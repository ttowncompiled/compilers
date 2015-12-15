## Massaged Productions (Final) - Ian Riley

1 _program_ := **program id (** <_identifier-list_> **) ;** <_program-body_>

1.1 _program-body_ := <_declarations_> <_program-subbody_> <br>
1.1 _program-body_ := <_program-subbody_>

1.2 _program-subbody_ := <_subprogram-declarations_> <_compound-statement_> **.** <br>
1.2 _program-subbody_ := <_compound-statement_> **.**

3 _identifier-list_ := **id** <_identifier-list'_>

3.1 _identifier-list'_ := **, id** <_identifier-list'_> <br>
3.1 _identifier-list'_ := **epsilon**

4 _declarations_ := **var id :** <_type_> **;** <_declarations'_>

4.1 _declarations'_ := **var id :** <_type_> **;** <_declarations'_> <br>
4.1 _declarations'_ := **epsilon**

5 _type_ := <_standard-type_> <br>
5 _type_ := **array [ num .. num ] of** <_standard-type_>

6 _standard-type_ := **integer** <br>
6 _standard-type_ := **real**

7 _subprogram-declarations_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_>

7.1 _subprogram-declarations'_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_> <br>
7.1 _subprogram-declarations'_ := **epsilon**

8 _subprogram-declaration_ := <_subprogram-head_> <_subprogram-body_>

8.1 _subprogram-body_ := <_declarations_> <_subprogram-subbody_> <br>
8.1 _subprogram-body_ := <_subprogram-subbody_>

8.2 _subprogram-subbody_ := <_subprogram-declarations_> <_compound-statement_> <br>
8.2 _subprogram-subbody_ := <_compound-statement_>

9 _subprogram-head_ := **function id** <_subprogram-head'_>

9.1 _subprogram-head'_ := <_arguments_> **:** <_standard-type_> **;** <br>
9.1 _subprogram-head'_ := **:** <_standard-type_> **;**

10 _arguments_ := **(** <_parameter-list_> **)**

11 _parameter-list_ := **id :** <_type_> <_parameter-list'_>

11.1 _parameter-list'_ := **; id :** <_type_> <_parameter-list'_> <br>
11.1 _parameter-list'_ := **epsilon**

12 _compound-statement_ := **begin** <_compound-statement'_> _{ compound-statement.type := compound-statement'.type }_

12.1 _compound-statement'_ := <_optional-statements_> **end** _{ compound-statement'.type := optional-statements.type }_ <br>
12.1 _compound-statement'_ := **end** _{ compound-statement'.type := void }_

13 _optional-statements_ := <_statement-list_> _{ optional-statements.type := statement-list.type }_

14 _statement-list_ := <_statement_> <_statement-list'_> _{ statement-list.type := **if** statement.type = statement-list'.type = void **then** void **else** type-error }_

14.1 _statement-list'_ := **;** <_statement_> <_statement-list'_> _{ statement-list'.type := **if** statement.type = statement-list'.type = void **then** void **else** type-error }_ <br>
14.1 _statement-list'_ := **epsilon** _{ statement-list'.type := void }_

15 _statement_ := <_variable_> **assignop** <_expression_> _{ statement.type := **if** variable.type = expression.type **then** void **else** type-error* }_ <br>
15 _statement_ := <_compound-statement_> _{ statement.type := compound-statement.type }_ <br>
15 _statement_ := **if** <_expression_> **then** <_statement1_> <_statement'_> _{ statement.type := **if** expression.type = boolean **then** **if** statement1.type = statement'.type = void **then** void **else** type-error **else** type-error* }_ <br>
15 _statement_ := **while** <_expression_> **do** <_statement1_> _{ statement.type := **if** expression.type = boolean **then** statement1.type **else** type-error* }_

15.1 _statement'_ := **else** <_statement_> _{ statement'.type := statement.type }_ <br>
15.1 _statement'_ := **epsilon** _{ statement'.type := void }_

16 _variable_ := **id** _{ id.type := lookup(id.entry); variable'.in := id.type }_ <_variable'_> _{ variable.type := variable'.type }_

16.1 _variable'_ := **[** <_expression_> **]** _{ variable'.type := **if** expression.type = integer **and** variable'.in = array(s, t) **then** t **else** type-error* }_ <br>
16.1 _variable'_ := **epsilon** _{ variable'.type := variable'.in }_

17 _expression-list_ := <_expression_> <_expression-list'_>

17.1 _expression-list'_ := **,** <_expression_> <_expression-list'_> <br>
17.1 _expression-list'_ := **epsilon**

18 _expression_ := <_simple-expression_> _{ expression'.in := simple-expression.type }_ <_expression'_> _{ expression.type := expression'.type }_

18.1 _expression'_ := **relop** <_simple-expression_> _{ expression'.type := **if** expression'.in = simple-expression.type **then** boolean **else** type-error* }_ <br>
18.1 _expression'_ := **epsilon** _{ expression'.type := expression'.in }_

19 _simple-expression_ := <_term_> _{ simple-expression'.in := term.type }_ <_simple-expression'_> _{ simple-expression.type := simple-expression'.type }_ <br>
19 _simple-expression_ := <_sign_> <_term_> _{ simple-expression'.in := term.type }_ <_simple-expression'_> _{ simple-expression.type := simple-expression'.type }_

19.1 _simple-expression'_ := **addop** <_term_> _{ simple-expression1'.in := **if** simple-expression'.in = term.type **then** term.type **else** type-error* }_ <_simple-expression1'_> _{ simple-expression'.type := simple-expression1'.type }_ <br>
19.1 _simple-expression'_ := **epsilon** _{ simple-expression'.type := simple-expression'.in }_

20 _term_ := <_factor_> <_term'_>

20.1 _term'_ := **mulop** _{factor.in := term'.type}_ <_factor_> _{}_ <_term'_> <br>
20.1 _term'_ := **epsilon**

21 _factor_ := **id** _{factor'.in := id.type}_ <_factor'_> _{factor.type := factor'.type}_<br>
21 _factor_ := **num** _{factor.type := num.type}_ <br>
21 _factor_ := **(** <_expression_> **)** _{factor.type := expression.type}_ <br>
21 _factor_ := **not** <_factor1_> _{factor1.type == BOOLEAN; factor.type := factor1.type}_ 

21.1 _factor'_ := _{factor'.in.type == FUNCTION}_ **(** _{expression-list.in := factor'.in.params}_ <_expression-list_> **)** _{factor'.type := factor'.in.return.type}_<br>
21.1 _factor'_ := _{factor'.in.type == ARRAY}_ **[** _{expression.in := INT}_ <_expression_> **]** _{factor'.type := factor'.in.val.type}_<br>
21.1 _factor'_ := **epsilon** _{factor'.type := factor'.in.type}_

22 _sign_ := **+** <br>
22 _sign_ := **-**
