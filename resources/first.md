## First - Ian Riley

01.00 f(_program_) := {
  **program** (1)
}

01.01 f(_program-body_) := {
  f(_declarations_) (1),
  f(_program-subbody_) (2)
}

01.02 f(_program-subbody_) := {
  f(_subprogram-declarations_) (1),
  f(_compound-statement_) (2)
}

02.00 f(_id_) := {
  **id** (1)
}

03.00 f(_identifier-list_) := {
  **id** (1)
}

03.01 f(_identifier-list'_) := {
  **,** (1),
  **epsilon** (2)
}

04.00 f(_declarations_) := {
  **var** (1)
}

04.01 f(_declarations'_) := {
  **var** (1),
  **epsilon** (2)
}

05.00 f(_type_) := {
  f(_standard-type_) (1),
  **array** (2)
}

06.00 f(_standard-type_) := {
  **integer** (1),
  **real** (2)
}

07.00 f(_subprogram-declarations_) := {
  f(_subprogram-declaration_) (1)
}

07.01 f(_subprogram-declarations'_) := {
  f(_subprogram-declaration_) (1),
  **epsilon** (2)
}

08.00 f(_subprogram-declaration_) := {
  f(_subprogram-head_) (1)
}

08.01 f(_subprogram-body_) := {
  f(_declarations_) (1),
  f(_subprogram-subbody_) (2)
}

08.02 f(_subprogram-subbody_) := {
  f(_subprogram-declarations_) (1),
  f(_compound-statement_) (2)
}

09.00 f(_subprogram-head_) := {
  **function** (1)
}

09.01 f(_subprogram-head'_) := {
  f(_arguments_) (1),
  **:** (2)
}

10.00 f(_arguments_) := {
  **(** (1)
}

11.00 f(_parameter-list_) := {
  f(_id_) (1)
}

11.01 f(_parameter-list'_) := {
  **;** (1),
  **epsilon** (2)
}

12.00 f(_compound-statement_) := {
  **begin** (1)
}

12.01 f(_compound-statement'_) := {
  f(_optional-statements_) (1),
  **end** (2)
}

13.00 f(_optional-statements_) := {
  f(_statement-list_) (1)
}

14.00 f(_statement-list_) := {
  f(_statement_) (1)
}

14.01 f(_statement-list'_) := {
  **;** (1),
  **epsilon** (2)
}

15.00 f(_statement_) := {
  f(_variable_) (1),
  f(_compound-statement_) (2),
  **if** (3),
  **while** (4)
}

15.01 f(_statement'_) := {
  **else** (1),
  **epsilon** (2)
}

16.00 f(_variable_) := {
  **id** (1)
}

16.01 f(_variable'_) := {
  **[** (1),
  **epsilon** (2)
}

17.00 f(_expression-list_) := {
  f(_expression_) (1)
}

17.01 f(_expression-list'_) := {
  **,** (1),
  **epsilon** (2)
}

18.00 f(_expression_) := {
  f(_simple-expression_) (1)
}

18.01 f(_expression'_) := {
  **relop** (1),
  **epsilon** (2)
}

19.00 f(_simple-expression_) := {
  f(_term_) (1),
  f(_sign_) (2)
}

19.01 f(_simple-expression'_) := {
  **addop** (1),
  **epsilon** (2)
}

20.00 f(_term_) := {
  f(_factor_) (1)
}

20.01 f(_term'_) := {
  **mulop** (1),
  **epsilon** (2)
}

21.00 f(_factor_) := {
  **id** (1),
  **num** (2),
  **(** (3),
  **not** (4)
}

21.01 f(_factor'_) := {
  **(** (1),
  **[** (2),
  **epsilon** (3)
}

22.00 f(_sign_) := {
  **+** (1),
  **-** (2)
}
