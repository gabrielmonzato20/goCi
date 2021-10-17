package main

func TestSoma(t *testing.T){
  total := Soma(10,100)
  if total != 110{
t.Errorf("Resultado da soma invalid %d , Expect %d",total , 110)
  }

}
