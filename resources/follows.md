## Follows - Ian Riley

01.00 F(_program_) := {
  **$**
}

01.01 F(_program-body_) := {
  **$**
}

01.02 F(_program-subbody_) := {
  **$**
}

02.00 F(_id_) := {
  **:**
}

03.00 F(_identifier-list_) := {
  **)**
}

03.01 F(_identifier-list'_) := {
  **)**
}

04.00 F(_declarations_) := {
  **function**,
  **begin**
}

04.01 F(_declarations'_) := {
  **function**,
  **begin**
}

05.00 F(_type_) := {
  **;**,
  **)**
}

06.00 F(_standard-type_) := {
  **;**,
  **)**
}

07.00 F(_subprogram-declarations_) := {
  **begin**
}

07.01 F(_subprogram-declarations'_) := {
  **begin**
}

08.00 F(_subprogram-declaration_) := {
  **;**
}

08.01 F(_subprogram-body_) := {
  **;**
}

08.02 F(_subprogram-subbody_) := {
  **;**
}

09.00 F(_subprogram-head_) := {
  **var**,
  **function**,
  **begin**
}

09.01 F(_subprogram-head'_) := {
  **var**,
  **function**,
  **begin**
}

10.00 F(_arguments_) := {
  **:**
}

11.00 F(_parameter-list_) := {
  **)**
}

11.01 F(_parameter-list'_) := {
  **)**
}

12.00 F(_compound-statement_) := {
  **.**,
  **;**,
  **end**,
  **else**
}

12.01 F(_compound-statement'_) := {
  **.**,
  **;**,
  **end**
  **else**
}

13.00 F(_optional-statements_) := {
  **end**
}

14.00 F(_statement-list_) := {
  **end**
}

14.01 F(_statement-list'_) := {
  **end**
}

15.00 F(_statement_) := {
  **;**,
  **end**,
  **else**
}

15.01 F(_statement'_) := {
  **;**,
  **end**,
  **else**
}

16.00 F(_variable_) := {
  **assignop**
}

16.01 F(_variable'_) := {
  **assignop**
}

17.00 F(_expression-list_) := {
  **)**
}

17.01 F(_expression-list'_) := {
  **)**
}

18.00 F(_expression_) := {
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**
  **)**
}

18.01 F(_expression'_) := {
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**
  **)**
}

19.00 F(_simple-expression_) := {
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

19.01 F(_simple-expression'_) := {
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

20.00 F(_term_) := {
  **addop**,
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

20.01 F(_term'_) := {
  **addop**,
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

21.00 F(_factor_) := {
  **mulop**,
  **addop**,
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

21.01 F(_factor'_) := {
  **mulop**,
  **addop**,
  **relop**,
  **;**,
  **end**,
  **else**,
  **then**,
  **do**,
  **]**,
  **,**,
  **)**
}

22.00 F(_sign_) := {
  **id**,
  **num**,
  **(**,
  **not**
}
