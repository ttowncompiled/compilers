## Follows - Ian Riley

01.00 F(_program_) := {
  **$**
}

01.01 F(_program-body_) := {
  F(_program_)
}

01.02 F(_program-subbody_) := {
  F(_program-body)
}

02.00 F(_id_) := {
  **:**
}

03.00 F(_identifier-list_) := {
  **)**
}

03.01 F(_identifier-list'_) := {
  F(_identifier-list_)
}

04.00 F(_declarations_) := {
  f(_program-subbody_),
  f(_subprogram-subbody_)
}

04.01 F(_declarations'_) := {
  F(_declarations_)
}

05.00 F(_type_) := {
  **;**,
  f(_parameter-list'_)
}

06.00 F(_standard-type_) := {
  F(_type_),
  **;**
}

07.00 F(_subprogram-declarations_) := {
  f(_compound-statement_)
}

07.01 F(_subprogram-declarations'_) := {
  F(_subprogram-declarations_)
}

08.00 F(_subprogram-declaration_) := {
  **;**
}

08.01 F(_subprogram-body_) := {
  F(_subprogram-declaration_)
}

08.02 F(_subprogram-subbody_) := {
  F(_subprogram-body_)
}

09.00 F(_subprogram-head_) := {
  f(_subprogram-body_)
}

09.01 F(_subprogram-head'_) := {
  F(_subprogram-head_)
}

10.00 F(_arguments_) := {
  **:**
}

11.00 F(_parameter-list_) := {
  **)**
}

11.01 F(_parameter-list'_) := {
  F(_parameter-list_)
}

12.00 F(_compound-statement_) := {
  **.**,
  F(_subprogram-subbody_),
  F(_statement_)
}

12.01 F(_compound-statement'_) := {
  F(_compound-statement_)
}

13.00 F(_optional-statements_) := {
  **end**
}

14.00 F(_statement-list_) := {
  F(_optional-statements_)
}

14.01 F(_statement-list'_) := {
  F(_statement-list_)
}

15.00 F(_statement_) := {
  f(_statement-list'_),
  f(_statement'_)
}

15.01 F(_statement'_) := {
  F(_statement_)
}

16.00 F(_variable_) := {
  **assignop**
}

16.01 F(_variable'_) := {
  F(_variable_)
}

17.00 F(_expression-list_) := {
  **)**
}

17.01 F(_expression-list'_) := {
  F(_expression-list_)
}

18.00 F(_expression_) := {
  F(_statement_),
  **then**,
  **do**,
  **]**,
  f(_expression-list'_),
  **)**
}

18.01 F(_expression'_) := {
  F(_expression_)
}

19.00 F(_simple-expression_) := {
  f(_expression'_),
  F(_expression'_)
}

19.01 F(_simple-expression'_) := {
  F(_simple-expression_)
}

20.00 F(_term_) := {
  f(_simple-expression'_)
}

20.01 F(_term'_) := {
  F(_term_)
}

21.00 F(_factor_) := {
  f(_term'_)
}

21.01 F(_factor'_) := {
  F(_factor_)
}

22.00 F(_sign_) := {
  f(_term_)
}
