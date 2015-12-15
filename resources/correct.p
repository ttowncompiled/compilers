program example(input, output);
var a : integer;
var b: integer;
var d : real;
var n: array[1..10] of integer;
var o: array[1..10] of real;
function gcd(a: integer; b: integer): integer;
begin
  a := 1;
  d := 1.2;
  d := 1.2E2;
  if a < 1 then b := -1;
  if a > 1 then b := +2 else b := 1;
  while b <> 3 do b := b + 1;
  if b = 3 then b := 0;
  a := a + 1;
  a := a - 1;
  b := a or b;
  b := a mod 2;
  b := b div d;
  d := d * 2.0E1;
  d := d / 2.0E1;
  e := e and not 1;
  n[gcd(a, b)] := 1
end;
begin
end.
