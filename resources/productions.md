## Productions - Ian Riley

01.00 _program_ :=
  1. **program id (** _identifier-list_ **) ;**
    - _program-body_

01.01 _program-body_ :=
  1. _declarations_
    - _program-subbody_
  2. _program-subbody_

01.02 _program-subbody_ :=
  1. _subprogram-declarations_
    - _compound-statement_
    - **.**
  2. _compound-statement_
    - **.**

02.00 _id_ :=
  1. **id**

03.00 _identifier-list_ :=
  1. **id** _identifier-list'_

03.01 _identifier-list'_ :=
  1. **, id** _identifier-list'_
  2. **epsilon**

04.00 _declarations_ :=
  1. **var** _id_ **:** _type_ **;** _declarations'_

04.01 _declarations'_ :=
  1. **var** _id_ **:** _type_ **;** _declarations'_
  2. **epsilon**

05.00 _type_ :=
  1. _standard-type_
  2. **array [ num .. num ] of** _standard-type_

06.00 _standard-type_ :=
  1. **integer**
  2. **real**

07.00 _subprogram-declarations_ :=
  1. _subprogram-declaration_ **;** _subprogram-declarations'_

07.01 _subprogram-declarations'_ :=
  1. _subprogram-declaration_ **;** _subprogram-declarations'_
  2. **epsilon**

08.00 _subprogram-declaration_ :=
  1. _subprogram-head_
    - _subprogram-body_

08.01 _subprogram-body_ :=
  1. _declarations_
    - _subprogram-subbody_
  2. _subprogram-subbody_

08.02 _subprogram-subbody_ :=
  1. _subprogram-declarations_
    - _compound-statement_
  2. _compound-statement_

09.00 _subprogram-head_ :=
  1. **function id** _subprogram-head'_

09.01 _subprogram-head'_ :=
  1. _arguments_ **:** _standard-type_ **;**
  2. **:** _standard-type_ **;**

10.00 _arguments_ :=
  1. **(** _parameter-list_ **)**

11.00 _parameter-list_ :=
  1. _id_ **:** _type_ _parameter-list'_

11.01 _parameter-list'_ :=
  1. **;** _id_ **:** _type_ _parameter-list'_
  2. **epsilon**

12.00 _compound-statement_ :=
  1. **begin**
    - _compound-statement'_

12.01 _compound-statement'_ :=
  1. _optional-statements_
    - **end**
  2. **end**

13.00 _optional-statements_ :=
  1. _statement-list_

14.00 _statement-list_ :=
  1. _statement_ _statement-list'_

14.01 _statement-list'_ :=
  1. **;** _statement_ _statement-list'_
  2. **epsilon**

15.00 _statement_ :=
  1. _variable_ **assignop** _expression_
  2. _compound-statement_
  3. **if** _expression_ **then** _statement_ _statement'_
  4. **while** _expression_ **do** _statement_

15.01 _statement'_ :=
  1. **else** _statement_
  2. **epsilon**

16.00 _variable_ :=
  1. **id** _variable'_

16.01 _variable'_ :=
  1. **[** _expression_ **]**
  2. **epsilon**

17.00 _expression-list_ :=
  1. _expression_ _expression-list'_

17.01 _expression-list'_ :=
  1. **,** _expression_ _expression-list'_
  2. **epsilon**

18.00 _expression_ :=
  1. _simple-expression_ _expression'_

18.01 _expression'_ :=
  1. **relop** _simple-expression_
  2. **epsilon**

19.00 _simple-expression_ :=
  1. _term_ _simple-expression'_
  2. _sign_ _term_ _simple-expression'_

19.01 _simple-expression'_ :=
  1. **addop** _term_ _simple-expression'_
  2. **epsilon**

20.00 _term_ :=
  1. _factor_ _term'_

20.01 _term'_ :=
  1. **mulop** _factor_ _term'_
  2. **epsilon**

21.00 _factor_ :=
  1. **id** _factor'_
  2. **num**
  3. **(** _expression_ **)**
  4. **not** _factor_

21.01 _factor'_ :=
  1. **(** _expression-list_ **)**
  2. **[** _expression_ **]**
  3. **epsilon**

22.00 _sign_ :=
  1. **+** 
  2. **-**
