package main

import "testing"

func TestSoma(t *testing.T){
  total := soma(10,100)
  if total != 110{
t.Errorf("Result of sum is  invalid %d , Expect %d",total , 110)
  }

}
