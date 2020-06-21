1. NFA 消空边细节
2. 无空边NFA 如何 转 DFA
3. 逻辑单元的界定 (如TCP、帧、IP数据报....)
4. 命名的规定 (规定一些实体命名，这么在取名时就不会很难受了)
5. 类的共享变量的界定(readingPosition..有的对象内部可能还有数组，这时还要弄个readingPosition吗？)
6.
    ```go
    func (stf *StateTableFormer)sentenceIsBlank(sentence *sentence) bool{
        return len(sentence.symbols)==1 && sentence.symbols[0]==blankSymbol
    }
    
    ```
    这个判断应该放在stateTableFormer好，还是放在sentence好呢？
    
7. 关于配置
    `(grammarConf *conf.GrammarConf,lexicalConf *conf.LexicalConf)`  -> 不太好
    `(cf *conf.Conf)` -> 这个更好
    

8. 全局变量要慎用

9. 尽量用 NewProduction 这样的函数初始化对象，不要用
            &production{
                blankSymbol:stf.blankSymbol,
                additionCharBeginChar: 'a',
            }
    这样很难差错！！！
10. 能通过配置创建的对象，不要作为函数参数传入
    ```go
    // 这个是较好
    func function(conf *conf.Conf){
        item := &isMatchOfNFATestItem{
            nfaBuilder:machine.NewNFABuilder(conf),
        }
    }
    
    //这个不好
    func function(nfaBuilder *NFABuilder) {
        item := &isMatchOfNFATestItem{
            nfaBuilder: nfaBuilder,
        }
    }
    ```
11. 配置对象 不要出现在成员变量中
12. 配置对象的字段 是否应该作为成员变量，还是随意引用呢？
    ```go
        type sentence struct{
            symbols []string
            delimiterOfSymbols string
        }
        
        func NewSentence(symbols []string,cf *conf.Conf) *sentence{
            return &sentence{
                symbols:symbols,
                delimiterOfSymbols:cf.SyntaxConf.DelimiterOfSymbols,
            }
        }
        func (s *sentence)Parse(line string) {
            line = strings.TrimSpace(line)
            s.symbols = strings.Split(line,s.delimiterOfSymbols)
        }
   
    ```
    `delimiterOfSymbols` 是否有必要出现在sentence的字段中呢？
    还是像下面这样更好呢？
     ```go
            type sentence struct{
                symbols []string
            }
            
            func NewSentence(symbols []string,cf *conf.Conf) *sentence{
                return &sentence{
                    symbols:symbols,
                }
            }
            func (s *sentence)Parse(line string) {
                line = strings.TrimSpace(line)
                // conf.GetConf() 可以获得全局配置对象
                s.symbols = strings.Split(line,conf.GetConf().delimiterOfSymbols)
            }
    
        ```
13. 职责分派很重要，需要进行思考，否则就会出现职责不明确、过多流浪数据的问题。
    编码时要注意迪米特法则
    
14. ```go
        func (tp *TokenParser) wordPairToToken(wordPair *machine.WordPair) *Token{
        	token := &Token{
        		wordPair.GetSpecialChar(),
        		tp.specialCharTable.GetCode(wordPair.GetSpecialChar(),wordPair.GetWord()),
        		tp.specialCharTable.GetType(wordPair.GetSpecialChar()),
        		wordPair.GetWord(),
        	}
        	return token
        }
    ```
    逻辑单元的转换，要放在上层
    如这，wordPairToToken，应该放在token蹭，而不应该放在machine层

16. 遇到长名称时，可以在降低语义的情况下，选择短名称
    如下:
    ```go
        leftNonTerminator := handlingProduction.leftNonTerminator
    ```
    之后用 leftNonTerminator 指代 handlingProduction.leftNonTerminator 就可以了
    
17.  不要用position代替index,(因为index语义更准确，且长度更短)
18. 不要出现中文！！！
19. 测试配置尽量独立出来
    可以创建一个新的配置文件，在测试的时候，用该配置文件初始化
    ```go
        // 在main中: conf.Init(configureFilePath)
        // 在测试中: conf.Init(testConfigureFilePath)
    ```
20. 全局配置嵌入会导致复用性变差...
21. sentenceIsBlank() 这个函数需要一个blankSymbol的变量，这样写时，上层
    的对象可以不用把blankSymbol传入sentence对象，而sentence.IsBlank(),此时
    sentence没有blankSymbol,那是应该把blankSymbol作为全局变量，还是把它作为成员字段
    还是将blankSymbol传入IsBlankSymbol呢？
    ```go
        // 简单描述: 
        // production拥有一个sentence字段，现在我们需要判断sentence是否为空，
        // 这个判断涉及到了一个blankSymbol。
        // 下面列出了 4 种写法..
        // 请问哪种写法好
    
        // 第一种情况,
        // blankSymbol 作为 production 的成员字段
        type production struct{
    	    sentence *sentence    
	        blankSymbol string
        }
        type sentence struct{
            symbols []string
        }

        func (p *production) sentenceIsBlank() bool{   
                return sentence.symbols[0]==p.blankSymbol
        }

        
        // 第二种情况
        // blankSymbol 作为 production 的成员字段
        // 调用时，将blankSymbol作为参数传入
         type production struct{
            sentence *sentence    
            blankSymbol string
        }
        type sentence struct{
            symbols []string
        }
        
        // 参数版
        // p调用时，采用 p.sentence.IsBlank(p.blankSymbol)
        func (s *sentence) IsBlank(blankSymbol string) bool{
    	        return s.symbols[0]==blankSymbol    
        }
        

        // 第三种情况
        // 将 blankSymbol 作为 sentence 的成员字段
        type production struct{
            sentence *sentence    
        }
        type sentence struct{
            symbols []string
            blankSymbol string
        }
        
        func (s *sentence) IsBlank() bool{
                return s.symbols[0]==s.blankSymbol
        }



        // 第四种情况
        // 将 blankSymbol 作为 全局常量。
        const global_blankSymbol = "XXX"
        type production struct{
            sentence *sentence    
        }
        type sentence struct{
            symbols []string
        }
        func (s *sentence) IsBlank() bool{
                return s.symbols[0]==global_blankSymbol
        }
    ``` 
    最终与光文哥的探讨下，我选择了第四种，因为production和sentence 应该是共享同一个blankSymbol。
    这样符合逻辑，虽然可能会带来复用性下降的问题。
    
22. 种别码不应自动编码...因为之后还用到种别码生成Symbol
    

