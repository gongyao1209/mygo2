package main

import (
	"fmt"
	"strings"
)

func main()  {

	s := GetArr()

	fmt.Println(s)

	return
	str := "gongyao"

	fmt.Println(str)

	fmt.Println(strings.Index(str, "n"))

	fmt.Println(strings.Split(str, "n"))
}

func GetArr() string {
	str := `15900249297
13512228620
18102123987
13652055528
13820294509
18230378188
18822298898
15222531669
18920639696
13602083942
15222342800
18299953278
13102055618
18678388868
13012226698
13820735262
18003194208
18902158008
13920217816
15222880318
13502122163
15100980499
18822708196
18633623573
15100226446
13633339587
18020041667
18222711200
18920118328
13853870864
13022270158
15624279333
13312050368
13920585944
13821170396
18322248333
17073550889
18920773791
13820083925
18920038858
18822479227
18643270687
13920487112
18920619188
13752477011
18722627711
18622096376
15902263722
15030882207
13662039006
13703320816
17720039192
13082001366
18920289656
13672018865
15222158558
18222307711
18920753117`

	mobiles := strings.Split(str, "\n")

	newM := make([]string, 0, 10)
	for _, m := range mobiles {
		temp := fmt.Sprintf("'%s'", m)
		newM = append(newM, temp)
	}

	return strings.Join(newM, ",")
}