## First - Ian Riley

01.00 f(_program_) := {
  **program** (1)
}

01.01 f(_program-body_) := {
  **var** (1),
  **function** (2),
  **begin** (2)
}

01.02 f(_program-subbody_) := {
  **function** (1),
  **begin** (2)
}

02.00 f(_id_) := {
  **id** (1)
}

03.00 f(_identifier-list_) := {
  **id** (1)
}

03.01 f(_identifier-list'_) := {
  **,** (1),
  **)** (2)
}

04.00 f(_declarations_) := {
  **var** (1)
}

04.01 f(_declarations'_) := {
  **var** (1),
  **function** (2),
  **begin** (2)
}

05.00 f(_type_) := {
  **integer** (1),
  **real** (1)
  **array** (2)
}

06.00 f(_standard-type_) := {
  **integer** (1),
  **real** (2)
}

07.00 f(_subprogram-declarations_) := {
  **function** (1)
}

07.01 f(_subprogram-declarations'_) := {
  **function** (1),
  **begin** (2)
}

08.00 f(_subprogram-declaration_) := {
  **function** (1)
}

08.01 f(_subprogram-body_) := {
  **var** (1),
  **function** (2),
  **begin** (2)
}

08.02 f(_subprogram-subbody_) := {
  **function** (1),
  **begin** (2)
}

09.00 f(_subprogram-head_) := {
  **function** (1)
}

09.01 f(_subprogram-head'_) := {
  **(** (1),
  **:** (2)
}

10.00 f(_arguments_) := {
  **(** (1)
}

11.00 f(_parameter-list_) := {
  **id** (1)
}

11.01 f(_parameter-list'_) := {
  **;** (1),
  **)** (2)
}

12.00 f(_compound-statement_) := {
  **begin** (1)
}

12.01 f(_compound-statement'_) := {
  **id** (1),
  **begin** (1),
  **if** (1),
  **while** (1),
  **end** (2)
}

13.00 f(_optional-statements_) := {
  **id** (1),
  **begin** (1),
  **if** (1),
  **while** (1)
}

14.00 f(_statement-list_) := {
  **id** (1),
  **begin** (1),
  **if** (1),
  **while** (1)
}

14.01 f(_statement-list'_) := {
  **;** (1),
  **end** (2)
}

15.00 f(_statement_) := {
  **id** (1),
  **begin** (2),
  **if** (3),
  **while** (4)
}

15.01 f(_statement'_) := {
  **else** (1),
  **;** (2),
  **end** (2)
}

16.00 f(_variable_) := {
  **id** (1)
}

16.01 f(_variable'_) := {
  **[** (1),
  **assignop** (2)
}

17.00 f(_expression-list_) := {
  **id** (1),
  **num** (1),
  **(** (1),
  **not** (1),
  **+** (1),
  **-** (1)
}

17.01 f(_expression-list'_) := {
  **,** (1),
  **)** (2)
}

18.00 f(_expression_) := {
  **id** (1),
  **num** (1),
  **(** (1),
  **not** (1),
  **+** (1),
  **-** (1)
}

18.01 f(_expression'_) := {
  **relop** (1),
  **else** (2),
  **;** (2),
  **end** (2),
  **then** (2),
  **do** (2),
  **]** (2),
  **)** (2)
}

19.00 f(_simple-expression_) := {
  **id** (1),
  **num** (1),
  **(** (1),
  **not** (1),
  **+** (2),
  **-** (2)
}

19.01 f(_simple-expression'_) := {
  **addop** (1),
  **relop** (2),
  **else** (2),
  **;** (2),
  **end** (2),
  **then** (2),
  **do** (2),
  **]** (2),
  **)** (2)
}

20.00 f(_term_) := {
  **id** (1),
  **num** (1),
  **(** (1),
  **not** (1)
}

20.01 f(_term'_) := {
  **mulop** (1),
  **addop** (2),
  **relop** (2),
  **else** (2),
  **;** (2),
  **end** (2),
  **then** (2),
  **do** (2),
  **]** (2),
  **)** (2)
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
  **mulop** (3),
  **addop** (3),
  **relop** (3),
  **else** (3),
  **;** (3),
  **end** (3),
  **then** (3),
  **do** (3),
  **]** (3),
  **)** (3)
}

22.00 f(_sign_) := {
  **+** (1),
  **-** (2)
}
