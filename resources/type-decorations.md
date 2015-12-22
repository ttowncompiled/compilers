## Type Decorations - Ian Riley

1 _program_ := **program id** _{ addType(id.lex, PROGRAM) }_ **(** <_identifier-list_> **) ;** <_program-body_>

1.1 _program-body_ := <_declarations_> <_program-subbody_> <br>
1.1 _program-body_ := <_program-subbody_>

1.2 _program-subbody_ := <_subprogram-declarations_> <_compound-statement_> **.** <br>
1.2 _program-subbody_ := <_compound-statement_> **.**

3 _identifier-list_ := **id** _{ addType(id.lex, PARG) }_ <_identifier-list'_>

3.1 _identifier-list'_ := **, id** _{ addType(id.lex, PARG) }_ <_identifier-list'_> <br>
3.1 _identifier-list'_ := **epsilon**

4 _declarations_ := **var id :** <_type_> _{ addType(id.lex, type.type) }_ **;** <_declarations'_>

4.1 _declarations'_ := **var id :** <_type_> _{ addType(id.lex, type.type) }_ **;** <_declarations'_> <br>
4.1 _declarations'_ := **epsilon**

5 _type_ := <_standard-type_> _{ type.type := standard-type.type }_ <br>
5 _type_ := **array [ num1 .. num2 ] of** <_standard-type_> _{ type.type := array(num2.lexval - num1.lexval, standard-type.type) }_

6 _standard-type_ := **integer** _{ standard-type.type := integer }_ <br>
6 _standard-type_ := **real** _{ standard-type.type := real }_

7 _subprogram-declarations_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_>

7.1 _subprogram-declarations'_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_> <br>
7.1 _subprogram-declarations'_ := **epsilon**

8 _subprogram-declaration_ := <_subprogram-head_> <_subprogram-body_>

8.1 _subprogram-body_ := <_declarations_> <_subprogram-subbody_> <br>
8.1 _subprogram-body_ := <_subprogram-subbody_>

8.2 _subprogram-subbody_ := <_subprogram-declarations_> <_compound-statement_> <br>
8.2 _subprogram-subbody_ := <_compound-statement_>

9 _subprogram-head_ := **function id** _{ addType(id.lex, void -> void) }_ <_subprogram-head'_> _{ modifyType(id.lex, subprogram-head'.type) }_

9.1 _subprogram-head'_ := <_arguments_> **:** <_standard-type_> **;** _{ subprogram-head'.type := arguments.type -> standard-type.type }_ <br>
9.1 _subprogram-head'_ := **:** <_standard-type_> **;** _{ subprogram-head'.type := void -> standard-type.type }_

10 _arguments_ := **(** <_parameter-list_> **)** _{ arguments.type := parameter-list.type }_

11 _parameter-list_ := **id :** <_type_> _{ addType(id.lex, type.type) }_ <_parameter-list'_> _{ parameter-list.type := type.type **concat** parameter-list'.type }_

11.1 _parameter-list'_ := **; id :** <_type_> _{ addType(id.lex, type.type) }_ <_parameter-list1'_> _{ parameter-list'.type := type.type **concat** parameter-list1'.type }_ <br>
11.1 _parameter-list'_ := **epsilon** _{ parameter-list'.type := void }_

12 _compound-statement_ := **begin** <_compound-statement'_> _{ compound-statement.type := compound-statement'.type }_

12.1 _compound-statement'_ := <_optional-statements_> **end** _{ compound-statement'.type := optional-statements.type }_ <br>
12.1 _compound-statement'_ := **end** _{ compound-statement'.type := void }_

13 _optional-statements_ := <_statement-list_> _{ optional-statements.type := statement-list.type }_

14 _statement-list_ := <_statement_> <_statement-list'_> _{ statement-list.type := **if** statement.type = statement-list'.type = void **then** void **else** type-error }_

14.1 _statement-list'_ := **;** <_statement_> <_statement-list'_> _{ statement-list'.type := **if** statement.type = statement-list'.type = void **then** void **else** type-error }_ <br>
14.1 _statement-list'_ := **epsilon** _{ statement-list'.type := void }_

15 _statement_ := <_variable_> **assignop** <_expression_> _{ statement.type := **if** variable.type = s -> t then if variable.type = expression.type **then** void **else** type-error* else if expression.type = t **then** void **else** type-error* }_ <br>
15 _statement_ := <_compound-statement_> _{ statement.type := compound-statement.type }_ <br>
15 _statement_ := **if** <_expression_> **then** <_statement1_> <_statement'_> _{ statement.type := **if** expression.type = boolean **then** **if** statement1.type = statement'.type = void **then** void **else** type-error **else** type-error* }_ <br>
15 _statement_ := **while** <_expression_> **do** <_statement1_> _{ statement.type := **if** expression.type = boolean **then** statement1.type **else** type-error* }_

15.1 _statement'_ := **else** <_statement_> _{ statement'.type := statement.type }_ <br>
15.1 _statement'_ := **epsilon** _{ statement'.type := void }_

16 _variable_ := **id** _{ variable'.in := getType(id.lex) }_ <_variable'_> _{ variable.type := variable'.type }_

16.1 _variable'_ := **[** <_expression_> **]** _{ variable'.type := **if** expression.type = integer **and** variable'.in = array(s, t) **then** t **else** type-error* }_ <br>
16.1 _variable'_ := **epsilon** _{ variable'.type := variable'.in }_

17 _expression-list_ := <_expression_> _{ expression-list'.in := **if** expression = expression-list.in[i] **then** expression-list.in **else** type-error* }_ <_expression-list'_> _{ expression-list.type := expression-list'.type }_

17.1 _expression-list'_ := **,** <_expression_> _{ expression-list1'.in := **if** expression = expression-list'.in[i] **then** expression-list'.in **else** type-error* }_ <_expression-list1'_> _{ expression-list'.type := expression-list1'.type }_ <br>
17.1 _expression-list'_ := **epsilon** _{ expression-list'.type := expression-list'.in }_

18 _expression_ := <_simple-expression_> _{ expression'.in := simple-expression.type }_ <_expression'_> _{ expression.type := expression'.type }_

18.1 _expression'_ := **relop** <_simple-expression_> _{ expression'.type := **if** expression'.in = simple-expression.type **then** boolean **else** type-error* }_ <br>
18.1 _expression'_ := **epsilon** _{ expression'.type := expression'.in }_

19 _simple-expression_ := <_term_> _{ simple-expression'.in := term.type }_ <_simple-expression'_> _{ simple-expression.type := simple-expression'.type }_ <br>
19 _simple-expression_ := <_sign_> <_term_> _{ simple-expression'.in := term.type }_ <_simple-expression'_> _{ simple-expression.type := simple-expression'.type }_

19.1 _simple-expression'_ := **addop** <_term_> _{ simple-expression1'.in := **if** simple-expression'.in = term.type **then** term.type **else** type-error* }_ <_simple-expression1'_> _{ simple-expression'.type := simple-expression1'.type }_ <br>
19.1 _simple-expression'_ := **epsilon** _{ simple-expression'.type := simple-expression'.in }_

20 _term_ := <_factor_> _{ term'.in := factor.type }_ <_term'_> _{ term.type := term'.type }_

20.1 _term'_ := **mulop** <_factor_> _{ term1'.in := **if** term'.in = factor.type **then** factor.type **else** type-error* }_ <_term1'_> _{ term'.type := term1'.type }_ <br>
20.1 _term'_ := **epsilon** _{ term'.type := term'.in }_

21 _factor_ := **id** _{ factor'.in := getType(id.lex) }_ <_factor'_> _{ factor.type := factor'.type }_ <br>
21 _factor_ := **num** _{ factor.type := num.type }_ <br>
21 _factor_ := **(** <_expression_> **)** _{ factor.type := expression.type }_ <br>
21 _factor_ := **not** <_factor1_> _{ factor.type := **if** factor1.type = boolean **then** boolean **else** type-error* }_ 

21.1 _factor'_ := **(** _{ expression-list.in := **if** factor'.in = s -> t **then** s **else** type-error }_ <_expression-list_> **)** _{ factor'.type := **if** expression-list.type = s **and** factor'.in = s -> t **then** t **else** type-error* }_ <br>
21.1 _factor'_ := **[** <_expression_> **]** _{ factor'.type := **if** expression.type = integer **and** factor'.in = array(s, t) **then** t **else** type-error* }_ <br>
21.1 _factor'_ := **epsilon** _{ factor'.type := factor'.in }_

22 _sign_ := **+** <br>
22 _sign_ := **-**
