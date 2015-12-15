## Massaged Productions (Removal of Immediate Left Recursion) - Ian Riley

1 _program_ := **program id (** <_identifier-list_> **) ;** <_declarations_> <_subprogram-declarations_> <_compound-statement_> **.** <br>
1 _program_ := **program id (** <_identifier-list_> **) ;** <_declarations_> <_compound-statement_> **.** <br>
1 _program_ := **program id (** <_identifier-list_> **) ;** <_subprogram-declarations_> <_compound-statement_> **.** <br>
1 _program_ := **program id (** <_identifier-list_> **) ;** <_compound-statement_> **.**

2 _declarations_ := **var id :** <_type_> **;** <_declarations'_>

2.1 _declarations'_ := **var id :** <_type_> **;** <_declarations'_> <br>
2.1 _declarations'_ := **epsilon**

3 _type_ := <_standard-type_> <br>
3 _type_ := **array [ num .. num ] of** <_standard-type_>

4 _standard-type_ := **integer** <br>
4 _standard-type_ := **real**

5 _subprogram-declarations_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_>

5.1 _subprogram-declarations'_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_> <br>
5.1 _subprogram-declarations'_ := **epsilon**

6 _subprogram-declaration_ := <_subprogram-head_> <_declarations_> <_subprogram-declarations_> <_compound-statement_> <br>
6 _subprogram-declaration_ := <_subprogram-head_> <_declarations_> <_compound-statement_> <br>
6 _subprogram-declaration_ := <_subprogram-head_> <_subprogram-declarations_> <_compound-statement_> <br>
6 _subprogram-declaration_ := <_subprogram-head_> <_compound-statement_>

7 _subprogram-head_ := **function id** <_arguments_> **:** <_standard-type_> **;** <br>
7 _subprogram-head_ := **function id :** <_standard-type_> **;**

8 _arguments_ := **(** <_parameter-list_> **)**

9 _parameter-list_ := **id :** <_type_> <br>
9 _parameter-list_ := <_parameter-list_> **; id :** <_type_>

10 _compound-statement_ := **begin** <_optional-statements_> **end** <br>
10 _compound-statement_ := **begin end**

11 _optional-statements_ := <_statement-list_>

12 _statement-list_ := <_statement_> <_statement-list'_>

12.1 _statement-list'_ := **;** <_statement_> <_statement-list'_> <br>
12.1 _statement-list'_ := **epsilon**

13 _statement_ := <_variable_> **assignop** <_expression_> <br>
13 _statement_ := <_procedure-statement_> <br>
13 _statement_ := <_compound-statement_> <br>
13 _statement_ := **if** <_expression_> **then** <_statement_> <br>
13 _statement_ := **if** <_expression_> **then** <_statement_> **else** <_statement_> <br>
13 _statement_ := **while** <_expression_> **do** <_statement_>

14 _variable_ := **id** <br>
14 _variable_ := **id [** <_expression_> **]**

15 _expression-list_ := <_expression_> <_expression-list'>

15.1 _expression-list'_ := **,** <_expression_> <_expression-list'_> <br>
15.1 _expression-list'_ := **epsilon**

16 _expression_ := <_simple-expression_> <br>
16 _expression_ := <_simple-expression_> **relop** <_simple-expression_>

17 _simple-expression_ := <_term_> <_simple-expression'_> <br>
17 _simple-expression_ := <_sign_> <_term_> <_simple-expression'_>

17.1 _simple-expression'_ := **addop** <_term_> <_simple-expression'_> <br>
17.1 _simple-expression'_ := **epsilon**

18 _term_ := <_factor_> <_term'_>

18.1 _term'_ := **mulop** <_factor_> <_term'_> <br>
18.1 _term'_ := **epsilon**

19 _factor_ := **id** <br>
19 _factor_ := **id (** <_expression-list_> **)** <br>
19 _factor_ := **id [** <_expression_> **]** <br>
19 _factor_ := **num** <br>
19 _factor_ := **(** <_expression_> **)** <br>
19 _factor_ := **not** <_factor_> <br>

20 _sign_ := **+** <br>
20 _sign_ := **-**
