package l1compiler

import ("fmt"; "os")

type Compiler struct {
  err string
  ast Node
  //codegen CodeGenerator
}


func NewCompiler() *Compiler {
  return &Compiler{}
}

/// Used in parser to set Root that everything will be attached to
func (c *Compiler) SetAstRoot(root Node) {
	c.ast = root
}

/// Performs Scan and Parsing on provided filepath and builds the AST in memory
/// Only valid for one program at a time. Once the compiler is used again,
/// the previous ast is removed
func (c *Compiler) parse(filename string) error {
  //TODO: Clear AST from the compiler object
  file, err := os.Open(filename)

  if err != nil {
    return err
  }

  // attaches Compiler pointer to the lexer so that it can be accessed from the
  // lexer and parser to save errors and the AST
  ret := yyParse(NewLexerWithInit(file, func (lex *Lexer) { lex.p = c}))

  if ret == 1 {
    return fmt.Errorf(filename + c.err)
  }

  return nil
}


/// Can only be called after a valid Parsing
// func (c *Compiler) generateCode() error {
//   c.codegen = L1toASMGenerator();
//   err := c.codegen.beginCompiler(c.ast)
//   return err
// }


/// Only exposed function that wraps the parse and code generation functions

func (c *Compiler) Compile(filename string) error {
  err := c.parse(filename)

  if err != nil {
    return err
  }

  // err = c.generateCode()
  //
  // if err != nil {
  //   fmt.Printf("Code generation error: %q%d", err)
  //   return err
  // }

  return nil
}
