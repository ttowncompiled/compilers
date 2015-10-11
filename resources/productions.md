## Productions - Ian Riley

01.00 _program_ :=
        - **program id (** _identifier-list_ **) ;**
        - _program-body_

01.01 _program-body_ :=
        - _declarations_
        - _program-subbody_
        - | _program-subbody_

01.02 _program-subbody_ :=
        - _subprogram-declarations_
        - _compound-statement_
        - **.**
        - | _compound-statement_
        - **.**

02.00 _id_ :=
        - **id**

03.00 _identifier-list_ :=
        - **id** _identifier-list'_

03.01 _identifier-list'_ :=
        - **, id** _identifier-list'_
        - | **epsilon**

04.00 _declarations_ :=
        - **var** _id_ **:** _type_ **;** _declarations'_

04.01 _declarations'_ :=
        - **var** _id_ **:** _type_ **;** _declarations'_
        - | **epsilon**

05.00 _type_ :=
        - _standard-type_
        - | **array [ num .. num ] of** _standard-type_

06.00 _standard-type_ :=
        - **integer**
        - | **real**

07.00 _subprogram-declarations_ :=
        - _subprogram-declaration_ **;** _subprogram-declarations'_

07.01 _subprogram-declarations'_ :=
        - _subprogram-declaration_ **;** _subprogram-declarations'_
        - | **epsilon**

08.00 _subprogram-declaration_ :=
        - _subprogram-head_
        - _subprogram-body_

08.01 _subprogram-body_ :=
        - _declarations_
        - _subprogram-subbody_
        - | _subprogram-subbody_

08.02 _subprogram-subbody_ :=
        - _subprogram-declarations_
        - _compound-statement_
        - | _compound-statement_

09.00 _subprogram-head_ :=
        - **function id** _subprogram-head'_

09.01 _subprogram-head'_ :=
        - _arguments_ **:** _standard-type_ **;**
        - | **:** _standard-type_ **;**

10.00 _arguments_ :=
        - **(** _parameter-list_ **)**

11.00 _parameter-list_ :=
        - _id_ **:** _type_ _parameter-list'_

11.01 _parameter-list'_ :=
        - **;** _id_ **:** _type_ _parameter-list'_
        - | **epsilon**

12.00 _compound-statement_ :=
        - **begin**
        - _compound-statement'_

12.01 _compound-statement'_ :=
        - _optional-statements_
        - **end**
        - | **end**

13.00 _optional-statements_ :=
        - _statement-list_

14.00 _statement-list_ :=
        - _statement_ _statement-list'_

14.01 _statement-list'_ :=
        - **;** _statement_ _statement-list'_
        - | **epsilon**

15.00 _statement_ :=
        - _variable_ **assignop** _expression_
        - | _compound-statement_
        - | **if** _expression_ **then** _statement_ _statement'_
        - | **while** _expression_ **do** _statement_

15.01 _statement'_ :=
        - **else** _statement_
        - | **epsilon**

16.00 _variable_ :=
        - **id** _variable'_

16.01 _variable'_ :=
        - **[** _expression_ **]**
        - | **epsilon**

17.00 _expression-list_ :=
        - _expression_ _expression-list'_

17.01 _expression-list'_ :=
        - **,** _expression_ _expression-list'_
        - | **epsilon**

18.00 _expression_ :=
        - _simple-expression_ _expression'_

18.01 _expression'_ :=
        - **relop** _simple-expression_
        - | **epsilon**

19.00 _simple-expression_ :=
        - _term_ _simple-expression'_
        - | _sign_ _term_ _simple-expression'_

19.01 _simple-expression'_ :=
        - **addop** _term_ _simple-expression'_
        - | **epsilon**

20.00 _term_ :=
        - _factor_ _term'_

20.01 _term'_ :=
        - **mulop** _factor_ _term'_
        - | **epsilon**

21.00 _factor_ :=
        - **id** _factor'_
        - | **num**
        - | **(** _expression_ **)**
        - | **not** _factor_

21.01 _factor'_ :=
        - **(** _expression-list_ **)**
        - | **[** _expression_ **]**
        - | **epsilon**

22.00 _sign_ :=
        - **+** | **-**
