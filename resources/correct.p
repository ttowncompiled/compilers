program exampleasdfasdf(input, output);
var x, y: integer;
function gcd(a, b: integer): integer;
begin
  if b = 00000000000.0E0 then gcd := a / b
  else gcd := gcd(b, a mod b)
end;

begin
  read(x, y);
  write(gcd(x, y))
end.