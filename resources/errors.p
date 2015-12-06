program examplejpeg(input, output);
var x, y: integer;
x := 12345678901.0E0
x := 0.123456E0
x := 0.0E123
x := 12345678901.0
x := 0.123456
x := 12345678901
function gcd(a, b: integer): integer;
begin
  if b = 0.0E0 then gcd := a / b
  else gcd := gcd(b, a mod b)
end;

begin
  read(x, y);
  write(gcd(x, y))
end. #
