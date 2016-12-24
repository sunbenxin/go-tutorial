# go-tutorial

## words to remember
 - qualified omitted explicit brace bracket encompasses extent

## go sytanx

### syntax

- ebnf

            |   alternation
            ()  grouping
            []  option (0 or 1 times)
            {}  repetition (0 to n times)

-  go programes are constructed by linking together packages.

         SourceFile   = PackageClause ";" { ImportDecl ";"} { TopLevelDecl ";" } .

         TopLevelDecl = Declaration | FunctionDecl | MethodDecl .

         Declaration  = Constecl | TypeDecl | VarDecl .

-  a package in turn is constructed from one or more source files that together declare constants,types,variables and 
    functions belonging to  the packages and which are accessible in all files of the same package.Those elemants may be
    exported and used in another package.
-  a package clause begins each source file and defines the package to which the file belongs.
-  PackageClause      = "package" PackageName. 
   PackageName         = identifier.
-  the package name must not be the blank identifier.
-  a set of files sharing the same PackageName form the implemention  of the package. A implemention may require that
    all source files for a package inhabit the same  directory.
-  ImportDecl         = "import" ( ImportSpec | "(" { ImportSpec } ")" ).
   ImportSpec         = [ "."| PackageName ] ImportPath
   ImportPath         = string_lit .
-  the PackageName in ImportSpec is used in qualified identifier to access exported identifiers of the package within the importing
     source file. It is declared in the file block. If the packagename is omitted, it defaults to the identifier specified in the package
     clause of the imported package. If an explicit period appears instead of name, all the package's exported identifiers declaried in that
     package block will be declared in the importing source file's file block and must be accessed without a qualifier.

### block

-  a block is possibly empty sequence of declarations and statements within matching brace brackets.
-  Block = "{" StatementList "}"
    StatementList = { Statement ";" } .
-  In addition to explicit blocks in the source code, there are implicit blocks:
    * The universe block encompasses all Go source text.
    * Each package  has a package block  containing  all Go source text for that package.
    * Each file has a file block containing all Go source text in that file.
    * Each "if", "for"  and "switch" statement is considered to be in  its own implicit block.
    * Each clause in a "switch" or "select" statement acts as an implicit block.

- a declaration binds a non-blank identifier to a constant,type,variable,function,label,or package.Every identifier in a program must be
    declared.No identifier may be declared twice in the same block,and no identifier may be declared in both the file and package block.
-  The blank identifier may be used like any other identifier in a declaration,but it does not  introduce a binding and thus is not declared.
    In the package block,the identifier init   may only be used for init function declarations, and  like the blank identifier it does not introuce
     a new binding.
-  the scope of a declared identifier is the extent of source text in which the identifier denotes the specified constant,type,variable,function
    label,or package.
     *   the predeclared identifier is the universe block.
     *   the scope of an identifier denoting a constant, type,variable,or function(but not method) declared at top level (outside any funciton) is the package block.
     *   the scope of the package name of an imported packages if the file block of the file containing the import declaration.
     *   the scope of an identifier denoting a method receiver,function parameter, or result variable is the funciion body.
     *   the scope of a constant or variable identifier declared inside a function begins at  the end of ConstSpec or VarSpec(ShortVarDecl for shor variable
     *    declarations) and ends at the end of the innermost containing block.
     *   the scope of a type identifier declared inside a function begins at the identifier in the TypeSpec and ends at the end of the innermost containing block.

-  an identifier declared in a block maybe reclared in an inner block.
-  the package clauses is not a declaration; the package name does not appear in any scope. Its purpose is to identify the files belonging to the
    same package and to specify the default package name for import declarations.
-  labels are declared  by labeled statements and are used in the "break" , "continue" and "goto" statements. It is  illegal to define a label that is never used.
       in contrast to other identifiers, labels are not block scoped  and do not conflict with identifiers that are not labels. The scope of a label
       is the function in which it is declared and exclude the body of any nested function.
-  the blank identifier is represented by the underscore character _.

### Program initialization and execution
- Within a package, package-level variables are initialized in declaration order but after any variables they depend on.
  a package-level variable is considered ready for initialized  if it is not yet initialized and either has no initialization expression or its initialization expression
  has no dependencies on uninitialized variables.Initializaiton proceeds by repeatedly  initializing the next package-level variable that is earliest in declaration order and ready for
  initialization,until there are no variables ready for initialization.

- If any variables are still uninitialized when this process ends,those variables are part of one or more initialization cycles,and the program is not valid.
- The declaration order of variables declared in multiple files is determined by the order in which the files are presented to the compiler.
- variables may also be initialized using functions named init declared in the package block.Mutiple such function may be defined.
- If multiple packages import a package, the imported package  will be initialized only once.
- Package initialization - variable initialization and the invocation of init functions - happends in a single goroutine, sequentially,one package at a time.
- An init function may launch other goroutines, which can run coucurrently with the initializaiton code.However,the 
    initializaiton always sequences the init functions:it will not invoke the next one until the previous has finished.
- to eusure reproducible initialization behavior,build systems are encoraged to present mutiple files belonging to the same package in lexical file name order to 
  to a compiler.

- a complete program is created by linking a single,unimported package called the main package with all the packages it imports,transitively. The main package must have package name
  main and declare a function main that takes no auguments and returns no value.

- Program execution begins by initializing the main package and then invoking the function main. When that function invocation returns,the program exits.It does not wait for other(non-main) 
  goroutines to complete.


- The predeclared type error is the conventional interface for representing an error condition,with the nil value representing no error.

- Execution errors such as attempting to index an array out of bounds trigger a run-time panic,equivalent to call of the built-in function panic with a value of th e
  implementation-definedinterface type runtime.Error.


### declarations
## constant decalrations
-  a constant declaration binds a list of identifiers to the values of a list of constant expressiongs.
    ConstDecl   =   "const" ( ConstSpec ) | "(" {  ConstSpec ";" } ")" ).
    ConstSpec   =   IdentifierList  [ [ Type ] "=" ExpressionList ].
    IdentifierList  =   identifier { "," identifier }.
    ExpressionList  = Expression { "," Expression }.

- If the type is present, all constants take the type specified.if the type if omitted,the constants
  take the individual types of the corresponding expressions.
  If the expression values are untyped constants,the declared constants remain untyped and the constant
  identifiers denote the constant values.

- within a parenthesized const declaration list the expression list may be omitted from any but the 
  first declaration. Such an empty list is equivalent to the textual substitution of the first preceding
  non-empty expression list and its type if any.Omitting the list of expressions is therefore equivalent to repeating
  the previous list.

- the predeclared identifier iota represents successive untyped integer constants.It is reset to 0 whenever the 
  reserved word const appears in the source and increments after each ConstSpec.

- within a ExpressionList, the value of each iota is the same because it is only incremented after each ConstSpec.

-
# Type declarations

- A type declaration binds an identifier, the type name, to a new type that has the same underlying type as an existing type.
  the operations defined for the existing tpe are also defined for the new type.
  The new type is different from the existing type.

    TypeDecl    =   "type" ( TypeSpec | "(" { TypeSpec ";" } ")" ).
    TypeSpec    =   identifier  Type .

- The declared type does not inherit any methods bound to the existing type, the the method set of an interface type
   or of emements of a composite type remains unchanged.

-# Variable declarations

- A variable declaration creates one or more variables, binds corresponding identifiers to them,and gives each a type and an initial value.

    VarDecl     =   "var" ( VarSpec | "(" { VarSpec ";" } ")" ).
    VarSpec     =   IdentifierList ( Type [ "=" ExpressionList ] | "=" ExpressionList )

-  If a type is not present,each variable is given the type of the corresponding initialization value in the assignment.
  if that value is an untyped constant,it is first converted to its default type. it it is an untyped boolean value, it is first converted to type bool.
    the predeclared value nil cannot be used to initialize a variable with no explicit type.

# unsafe
-  The built-in package unsafe,known to the compiler,provides facilities for low-level programming
	including operations that violate the type system. A pacakge using unsafe must be vetted 
	manually for type safety and may not be portable.

- built-in functions are predeclared.the built-in functions do not have standard Go types, so they 
	can only appear in call expressios;they cannot be used as function values.
-  for a channel c, the built-in function close(c) records that no more values will be sent on the
	channel.Tt is an error if c is a receive-only channel.Sending to or closing a closed channel	

###note
- golang function does not support function reload(函数重载)
