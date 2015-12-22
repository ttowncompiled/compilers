## Memory Decorations - Ian Riley

1 _program_ := **program id (** <_identifier-list_> **) ;** _{ loc := 0 }_ <_program-body_>

1.1 _program-body_ := <_declarations_> <_program-subbody_> <br>
1.1 _program-body_ := <_program-subbody_>

1.2 _program-subbody_ := <_subprogram-declarations_> <_compound-statement_> **.** <br>
1.2 _program-subbody_ := <_compound-statement_> **.**

3 _identifier-list_ := **id** <_identifier-list'_>

3.1 _identifier-list'_ := **, id** <_identifier-list'_> <br>
3.1 _identifier-list'_ := **epsilon**

4 _declarations_ := **var id :** <_type_> **;** _{ loc += computeMemory(type.type) }_ <_declarations'_>

4.1 _declarations'_ := **var id :** <_type_> **;** _{ loc += computeMemory(type.type) }_ <_declarations'_> <br>
4.1 _declarations'_ := **epsilon**

5 _type_ := <_standard-type_> <br>
5 _type_ := **array [ num .. num ] of** <_standard-type_>

6 _standard-type_ := **integer** <br>
6 _standard-type_ := **real**

7 _subprogram-declarations_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_>

7.1 _subprogram-declarations'_ := <_subprogram-declaration_> **;** <_subprogram-declarations'_> <br>
7.1 _subprogram-declarations'_ := **epsilon**

8 _subprogram-declaration_ := <_subprogram-head_> _{ offset := loc; loc := 0; }_ <_subprogram-body_> _{ loc := offset }_

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

12 _compound-statement_ := **begin** <_compound-statement'_>

12.1 _compound-statement'_ := <_optional-statements_> **end** <br>
12.1 _compound-statement'_ := **end**

13 _optional-statements_ := <_statement-list_>

14 _statement-list_ := <_statement_> <_statement-list'_>

14.1 _statement-list'_ := **;** <_statement_> <_statement-list'_> <br>
14.1 _statement-list'_ := **epsilon**

15 _statement_ := <_variable_> **assignop** <_expression_> <br>
15 _statement_ := <_compound-statement_> <br>
15 _statement_ := **if** <_expression_> **then** <_statement_> <_statement'_> <br>
15 _statement_ := **while** <_expression_> **do** <_statement_>

15.1 _statement'_ := **else** <_statement_> <br>
15.1 _statement'_ := **epsilon**

16 _variable_ := **id** <_variable'_>

16.1 _variable'_ := **[** <_expression_> **]** <br>
16.1 _variable'_ := **epsilon**

17 _expression-list_ := <_expression_> <_expression-list'_>

17.1 _expression-list'_ := **,** <_expression_> <_expression-list'_> <br>
17.1 _expression-list'_ := **epsilon**

18 _expression_ := <_simple-expression_> <_expression'_>

18.1 _expression'_ := **relop** <_simple-expression_> <br>
18.1 _expression'_ := **epsilon**

19 _simple-expression_ := <_term_> <_simple-expression'_> <br>
19 _simple-expression_ := <_sign_> <_term_> <_simple-expression'_>

19.1 _simple-expression'_ := **addop** <_term_> <_simple-expression'_> <br>
19.1 _simple-expression'_ := **epsilon**

20 _term_ := <_factor_> <_term'_>

20.1 _term'_ := **mulop** <_factor_> <_term'_> <br>
20.1 _term'_ := **epsilon**

21 _factor_ := **id** <_factor'_> <br>
21 _factor_ := **num** <br>
21 _factor_ := **(** <_expression_> **)** <br>
21 _factor_ := **not** <_factor_>

21.1 _factor'_ := **(** <_expression-list_> **)** <br>
21.1 _factor'_ := **[** <_expression_> **]** <br>
21.1 _factor'_ := **epsilon**

22 _sign_ := **+** <br>
22 _sign_ := **-**
