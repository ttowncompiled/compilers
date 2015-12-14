## Productions - Ian Riley

1 _program_ := **program id (** <_identifier-list_> **) ;** <_declarations_> <_subprogram-declarations_> <_compound-statement_> **.**

2 _declarations_ := <_declarations_> **var id :** <_type_> **;** <br>
2 _declarations_ := **epsilon**

3 _type_ := <_standard-type_> <br>
3 _type_ := **array [ num .. num ] of** <_standard-type_>

4 _standard-type_ := **integer** <br>
4 _standard-type_ := **real**

5 _subprogram-declarations_ := <_subprogram-declarations_> <_subprogram-declaration_> **;** <br>
5 _subprogram-declarations_ := **epsilon**

6 _subprogram-declaration_ := <_subprogram-head_> <_declarations_> <_subprogram-declarations_> <_compound-statement_>

7 _subprogram-head_ := **function id** <_arguments_> **:** <_standard-type_> **;**

8 _arguments_ := **(** <_parameter-list_> **)** <br>
8 _arguments_ := **epsilon**

9 _parameter-list_ := **id :** <_type_> <br>
9 _parameter-list_ := <_parameter-list_> **; id :** <_type_>

10 _compound-statement_ := **begin** <_optional-statements_> **end**

11 _optional-statements_ := <_statement-list_> <br>
11 _optional-statements_ := **epsilon**

12 _statement-list_ := <_statement_> <br>
12 _statement-list_ := <_statement-list_> **;** <_statement_>

13 _statement_ := <_variable_> **assignop** <_expression_> <br>
13 _statement_ := <_procedure-statement_> <br>
13 _statement_ := <_compound-statement_> <br>
13 _statement_ := **if** <_expression_> **then** <_statement_> <br>
13 _statement_ := **if** <_expression_> **then** <_statement_> **else** <_statement_> <br>
13 _statement_ := **while** <_expression_> **do** <_statement_>

14 _variable_ := **id** <br>
14 _variable_ := **id [** <_expression_> **]**

15 _expression-list_ := <_expression_> <br>
15 _expression-list_ := <_expression-list_> **,** <_expression_>

16 _expression_ := <_simple-expression_> <br>
16 _expression_ := <_simple-expression_> **relop** <_simple-expression_>

17 _simple-expression_ := <_term_> <br>
17 _simple-expression_ := <_sign_> <_term_> <br>
17 _simple-expression_ := <_simple-expression_> **addop** <_term_>

18 _term_ := <_factor_> <br>
18 _term_ := <_term_> **mulop** <_factor_>

19 _factor_ := **id** <br>
19 _factor_ := **id (** <_expression-list_> **)** <br>
19 _factor_ := **id [** <_expression_> **]** <br>
19 _factor_ := **num** <br>
19 _factor_ := **(** <_expression_> **)** <br>
19 _factor_ := **not** <_factor_> <br>

20 _sign_ := **+** <br>
20 _sign_ := **-**
