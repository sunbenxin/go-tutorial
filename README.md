# go-tutorial

## words to remember
 - qualified omitted explicit brace bracket encompasses extent

## go sytanx

### syntax

-  SourceFile         = PackageClause ";" { ImportDecl ";"} { TopLevelDecl ";" } .
         TopLevelDecl = Declaration | FunctionDecl | MethodDecl .
         Declaration = Constecl | TypeDecl | VarDecl .

-  go programes are constructed by linking together packages.
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
