01.00 _program_ :=
        - **program id (** _identifier-list_ **) ;**
        - _declarations_
        - _subprogram-declarations_
        - _compound-statement_
        - **.**
        - | **program id (** _identifier-list_ **) ;**
        - _declarations_
        - _compound-statement_
        - **.**
        - | **program id (** _identifier_list_ **) ;**
        - _subprogram-declarations_
        - _compound-statement_
        - **.**
        - | **program id (** _identifier_list_ **) ;**
        - _compound-statement_
        - **.**

02.00 _id_ :=
        - **id**

03.00 _declarations_ :=
        - **var** _id_ **:** _type_ **;** _declarations'_
        - | **var** _id_ **:** _type_ **;**

03.01 _declarations_ :=
        - **var** _id_ **:** _type_ **;** _declarations'_
        - | **var** _id_ **:** _type_ **;**

04.00 _type_ :=
        - _standard-type_
        - | **array [ num .. num ] of** _standard-type_

05.00 _standard-type_ :=
        - **integer**
        - | **real**

06.00 _subprogram-declarations_ :=
        - _subprogram-declaration_ **;** _subprogram-declarations'_
        - | _subprogram-declaration_ **;**

06.01 _subprogram-declarations'_ :=
        - _subprogram-declaration_ **;** _subprogram-declarations'_
        - | _subprogram-declaration_ **;**

07.00 _subprogram-declaration_ :=
        - _subprogram-head_
        - _declarations_
        - _subprogram-declarations_
        - _compound-statement_
        - | _subprogram-head_
        - _declarations_
        - _compound-statement_
        - | _subprogram-head_
        - _subprogram-declarations_
        - _compound-statement_
        - | _subprogram-head_
        - _compound-statement_

08.00 _subprogram-head_ :=
        - **function id** _arguments_ **:** _standard-type_ **;**
        - | **function id :** _standard-type_ **;**

09.00 _arguments_ :=
        - **(** _parameter-list_ **)**

10.00 _parameter-list_ :=
        - _id_ **:** _type_ _parameter-list'_
        - | _id_ **:** _type_

10.01 _parameter-list'_ :=
        - **;** _id_ **:** _type_ _parameter-list'_
        - | **;** _id_ **:** _type_

11.00 _compound-statement_ :=
        - **begin**
        - _optional-statements_
        - **end**
        - | **begin**
        - **end**

12.00 _optional-statements_ :=
        - _statement-list_

13.00 _statement-list_ :=
        - _statement_ _statement-list'_
        - | _statement_

13.01 _statement-list'_ :=
        - **;** _statement_ _statement-list'_
        - | **;** _statement_

14.00 _statement_ :=
        - _variable_ **assignop** _expression_
        - | _procedure-statement_
        - | _compound-statement_
        - | **if** _expression_ **then** _statement_
        - | **if** _expression_ **then** _statement_ **else** _statement_
        - | **while** _expression_ **do** _statement_

15.00 _variable_ :=
        - **id**
        - | **id [** _expression_ **]**

16.00 _expression-list_ :=
        - _expression_ _expression-list'_
        - | _expression_

16.01 _expression-list'_ :=
        - **,** _expression_ _expression-list'_
        - | **,** _expression_

17.00 _expression_ :=
        - _simple-expression_
        - | _simple-expression_ **relop** _simple-expression_

18.00 _simple-expression_ :=
        - _term_ _simple-expression'_
        - | _term_
        - | _sign_ _term_ _simple-expression'_
        - | _sign_ _term_

18.01 _simple-expression'_ :=
        - **addop** _term_ _simple-expression'_
        - | **addop** _term_

19.00 _term_ :=
        - _factor_ _term'_
        - | _factor_

19.01 _term'_ :=
        - **mulop** _factor_ _term'_
        - | **mulop** _factor_

20.00 _factor_ :=
        - **id**
        - | **id (** _expression-list_ **)**
        - | **id [** _expression_ **]**
        - | **num**
        - | **(** _expression_ **)**
        - | **not** _factor_

21.00 _sign_ :=
        - **+** | **-**
