## First - Ian Riley

01.00 f(_program_) := {
  **program**
}

01.01 f(_program-body_) := {
  f(_declarations_),
  f(_program-subbody_)
}

01.02 f(_program-subbody_) := {
  f(_subprogram-declarations_),
  f(_compound-statement_)
}

02.00 f(_id_) := {
  **id**
}

03.00 f(_identifier-list_) := {
  **id**
}

03.01 f(_identifier-list'_) := {
  **,**,
  **epsilon**
}

04.00 f(_declarations_) := {
  **var**
}

04.01 f(_declarations'_) := {
  **var**,
  **epsilon**
}

05.00 f(_type_) := {
  f(_standard-type_),
  **array**
}

06.00 f(_standard-type_) := {
  **integer**,
  **real**
}

07.00 f(_subprogram-declarations_) := {
  f(_subprogram-declaration_)
}

07.01 f(_subprogram-declarations'_) := {
  f(_subprogram-declaration_),
  **epsilon**
}

08.00 f(_subprogram-declaration_) := {
  f(_subprogram-head_)
}

08.01 f(_subprogram-body_) := {
  f(_declarations_),
  f(_subprogram-subbody_)
}

08.02 f(_subprogram-subbody_) := {
  f(_subprogram-declarations_),
  f(_compound-statement_)
}

09.00 f(_subprogram-head_) := {
  **function**
}

09.01 f(_subprogram-head'_) := {
  f(_arguments_),
  **:**
}

10.00 f(_arguments_) := {
  **(**
}

11.00 f(_parameter-list_) := {
  f(_id_)
}

11.01 f(_parameter-list'_) := {
  **;**,
  **epsilon**
}

12.00 f(_compound-statement_) := {
  **begin**
}

12.01 f(_compound-statement'_) := {
  f(_optional-statements_),
  **end**
}

13.00 f(_optional-statements_) := {
  f(_statement-list_)
}

14.00 f(_statement-list_) := {
  f(_statement_)
}

14.01 f(_statement-list'_) := {
  **;**,
  **epsilon**
}

15.00 f(_statement_) := {
  f(_variable_),
  f(_compound-statement_),
  **if**,
  **while**
}

15.01 f(_statement'_) := {
  **else**,
  **epsilon**
}

16.00 f(_variable_) := {
  **id**
}

16.01 f(_variable'_) := {
  **[**,
  **epsilon**
}

17.00 f(_expression-list_) := {
  f(_expression_)
}

17.01 f(_expression-list'_) := {
  **,**,
  **epsilon**
}

18.00 f(_expression_) := {
  f(_simple-expression_)
}

18.01 f(_expression'_) := {
  **relop**,
  **epsilon**
}

19.00 f(_simple-expression_) := {
  f(_term_),
  f(_sign_)
}

19.01 f(_simple-expression'_) := {
  **addop**,
  **epsilon**
}

20.00 f(_term_) := {
  f(_factor_)
}

20.01 f(_term'_) := {
  **mulop**,
  **epsilon**
}

21.00 f(_factor_) := {
  **id**,
  **num**,
  **(**,
  **not**
}

21.01 f(_factor'_) := {
  **(**,
  **[**,
  **epsilon**
}

22.00 f(_sign_) := {
  **+**,
  **-**
}
