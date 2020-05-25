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
