package fibo

//Count is getting fibonacchi number
func Count(num uint64) uint64 {

	mapArg := make(map[string]uint64) //мэп - где будет храниться два аргумента и конечный результат
	mapArg["firstArg"] = 0
	mapArg["secondArg"] = 0
	mapArg["result"] = 0

	return getNumber(num, mapArg)

}

func getNumber(index uint64, mapArg map[string]uint64) uint64 {

	if index == 0 { //если index = 0, выходим из рекурсии.
		return mapArg["result"]
	}

	if mapArg["result"] == 0 && mapArg["firstArg"] == 0 && mapArg["secondArg"] == 0 { //ловим первый проход.
		mapArg["secondArg"] = 1 //при первом проходе result = 0, firstArg = 0, secondArg = 1.
		return getNumber(index-1, mapArg)
	}

	mapArg["result"] = mapArg["firstArg"] + mapArg["secondArg"] //логика работы, пример: result = firstArg(5) + secondArg(3) = 8
	mapArg["secondArg"] = mapArg["firstArg"]                    //меняем местами. secondArg(3) = firstArg(5) , => secondArg = 5
	mapArg["firstArg"] = mapArg["result"]                       //firstArg(5) = result(8), firstArg => 8. На следующем шаге: result = firstArg(8)+ secondArg(5)

	return getNumber(index-1, mapArg)

}
