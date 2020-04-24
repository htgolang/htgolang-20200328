package main

import (
	"fmt"
	"strings"
)

var dreams string = `
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我我
	Five score years ago, a great American, in whose symbolic shadow we stand signed the Emancipation Proclamation. This momentous decree came as a great beacon light of hope to millions of Negro slaves who had been seared in the flames of withering injustice. It came as a joyous daybreak to end the long night of captivity.
　　But one hundred years later, we must face the tragic fact that the Negro is still not free. One hundred years later, the life of the Negro is still sadly crippled by the manacles of segregation and the chains of discrimination. One hundred years later, the Negro lives on a lonely island of poverty in the midst of a vast ocean of material prosperity. One hundred years later, the Negro is still languishing in the corners of American society and finds himself an exile in his own land. So we have come here today to dramatize an appalling condition.
　　In a sense we have come to our nation's capital to cash a check. When the architects of our republic wrote the magnificent words of the Constitution and the declaration of Independence, they were signing a promissory note to which every American was to fall heir. This note was a promise that all men would be guaranteed the inalienable rights of life, liberty, and the pursuit of happiness.
　　It is obvious today that America has defaulted on this promissory note insofar as her citizens of color are concerned. Instead of honoring this sacred obligation, America has given the Negro people a bad check which has come back marked "insufficient funds." But we refuse to believe that the bank of justice is bankrupt. We refuse to believe that there are insufficient funds in the great vaults of opportunity of this nation.
　　So we have come to cash this check -- a check that will give us upon demand the riches of freedom and the security of justice.
　　We have also come to this hallowed spot to remind America of the fierce urgency of now. This is no time to engage in the luxury of cooling off or to take the tranquilizing drug of gradualism. Now is the time to rise from the dark and desolate valley of segregation to the sunlit path of racial justice. Now is the time to open the doors of opportunity to all of God's children. Now is the time to lift our nation from the quicksands of racial injustice to the solid rock of brotherhood.
　	It would be fatal for the nation to overlook the urgency of the moment and to underestimate the determination of the Negro. This sweltering summer of the Negro's legitimate discontent will not pass until there is an invigorating autumn of freedom and equality. Nineteen sixty-three is not an end, but a beginning.
　　Those who hope that the Negro needed to blow off steam and will now be content will have a rude awakening if the nation returns to business as usual. There will be neither rest nor tranquility in America until the Negro is granted his citizenship rights. The whirlwinds of revolt will continue to shake the foundations of our nation until the bright day of justice emerges.
　　But there is something that I must say to my people who stand on the warm threshold which leads into the palace of justice. In the process of gaining our rightful place we must not be guilty of wrongful deeds. Let us not seek to satisfy our thirst for freedom by drinking from the cup of bitterness and hatred.
　　We must forever conduct our struggle on the high plane of dignity and discipline. We must not allow our creative protest to degenerate into physical violence. Again and again we must rise to the majestic heights of meeting physical force with soul force.
　　The marvelous new militancy which has engulfed the Negro community must not lead us to distrust of all white people, for many of our white brothers, as evidenced by their presence here today, have come to realize that their destiny is tied up with our destiny and their freedom is inextricably bound to our freedom. We cannot walk alone.
　　And as we walk, we must make the pledge that we shall march ahead. We cannot turn back. There are those who are asking the devotees of civil rights, "When will you be satisfied?" We can never be satisfied as long as our bodies, heavy with the fatigue of travel, cannot gain lodging in the motels of the highways and the hotels of the cities. We cannot be satisfied as long as the Negro's basic mobility is from a smaller ghetto to a larger one. We can never be satisfied as long as a Negro in Mississippi cannot vote and a Negro in New York believes he has nothing for which to vote. No, no, we are not satisfied, and we will not be satisfied until justice rolls down like waters and righteousness like a mighty stream.
　　I am not unmindful that some of you have come here out of great trials and tribulations. Some of you have come fresh from narrow cells. Some of you have come from areas where your quest for freedom left you battered by the storms of persecution and staggered by the winds of police brutality. You have been the veterans of creative suffering. Continue to work with the faith that unearned suffering is redemptive.
　　Go back to Mississippi, go back to Alabama, go back to Georgia, go back to Louisiana, go back to the slums and ghettos of our northern cities, knowing that somehow this situation can and will be changed. Let us not wallow in the valley of despair.
　　I say to you today, my friends, that in spite of the difficulties and frustrations of the moment, I still have a dream. It is a dream deeply rooted in the American dream.
　　I have a dream that one day this nation will rise up and live out the true meaning of its creed: "We hold these truths to be self-evident: that all men are created equal."
　　I have a dream that one day on the red hills of Georgia the sons of former slaves and the sons of former slaveowners will be able to sit down together at a table of brotherhood.
　　I have a dream that one day even the state of Mississippi, a desert state, sweltering with the heat of injustice and oppression, will be transformed into an oasis of freedom and justice.
　　I have a dream that my four children will one day live in a nation where they will not be judged by the color of their skin but by the content of their character.
　　I have a dream today.
　　I have a dream that one day the state of Alabama, whose governor's lips are presently dripping with the words of interposition and nullification, will be transformed into a situation where little black boys and black girls will be able to join hands with little white boys and white girls and walk together as sisters and brothers.
　　I have a dream today.
　　I have a dream that one day every valley shall be exalted, every hill and mountain shall be made low, the rough places will be made plain, and the crooked places will be made straight, and the glory of the Lord shall be revealed, and all flesh shall see it together.
　　This is our hope. This is the faith with which I return to the South. With this faith we will be able to hew out of the mountain of despair a stone of hope. With this faith we will be able to transform the jangling discords of our nation into a beautiful symphony of brotherhood. With this faith we will be able to work together, to pray together, to struggle together, to go to jail together, to stand up for freedom together, knowing that we will be free one day.
　　This will be the day when all of God's children will be able to sing with a new meaning, "My country, 'tis of thee, sweet land of liberty, of thee I sing. Land where my fathers died, land of the pilgrim's pride, from every mountainside, let freedom ring."
　　And if America is to be a great nation this must become true. So let freedom ring from the prodigious hilltops of New Hampshire.
　　Let freedom ring from the mighty mountains of New York.
　　Let freedom ring from the heightening Alleghenies of Pennsylvania!
　　Let freedom ring from the snowcapped Rockies of Colorado!
　　Let freedom ring from the curvaceous peaks of California!
　　But not only that; let freedom ring from Stone Mountain of Georgia!
　　Let freedom ring from Lookout Mountain of Tennessee!
　　Let freedom ring from every hill and every molehill of Mississippi. From every mountainside, let freedom ring.
　　When we let freedom ring, when we let it ring from every village and every hamlet, from every state and every city, we will be able to speed up that day when all of God's children, black men and white men, Jews and Gentiles, Protestants and Catholics, will be able to join hands and sing in the words of the old Negro spiritual, "Free at last! free at last! thank God Almighty, we are free at last!"
`
//2. "我有一个梦想" 中出现次数最多的top 10 字符，英文字母区分大小写
/*
为了提高通用性，可以使用rune数据类型来处理字符（当出现中文字符时也可以被正常处理）

注意：暂时不去除标点符号，文中可能出现中文标点符号和英文标点符号，这样的话不好处理，后续学了正则表达式后再优化。
 */

//1. 将字符串类型转换成run类型并存储到切片中
//定义一个rune类型空切片
var runeString []rune

//1. 先将字符串的空格和空行去除，然后将字符串类型转换成rune类型
func stringToRune(s string) []rune  {
	s = strings.Replace(s, " ", "",-1)
	s = strings.Replace(s, "\n", "", -1)
	runeString = []rune(s)
	return runeString
}

//2. 然后遍历rune类型切片，将rune类型切片内的所有元素作为key生成一个新的字典，元素出现的次数作为value
//定义一个空字典
var getMapFromRuneString = make(map[rune]int32)
func elementCount(r []rune) map[rune]int32  {
	for _, v := range runeString {
		getMapFromRuneString[v]++
	}
	return getMapFromRuneString
}

//3. 然后将键值对作为一个整体，放入一个空切片中作为一个元素,此时生成一个二维切片
var newRune [][]rune
func mapToSlice(myMap map[rune]int32) [][]rune {
	for k, v := range myMap {
		var fixed []rune //每次需要清空
		fixed = append(fixed, k)
		fixed = append(fixed, v)
		newRune = append(newRune, fixed)//生成一个二维切片,切片的每个元素就是k,v
	}
	/*
	返回的切片由map而来，每次得到的切片顺序不一定就相同，所以需要固定返回的切片元素顺序，最后交给冒泡算法来排序。
	不然并列的情况下，每次打印出的并列的元素并不一定相同
	以key的值来固定顺序
	 */
	return newRune
}

//4. 对二维切片排序，使用冒泡排序按key排序。升序。
func bubbleKey(f [][]rune) [][]rune{
	for j:=0;j<len(f)-1;j++{
		for k:=0;k<len(f)-j-1;k++{
			if f[k][0] > f[k+1][0] {
				tmp := f[k]
				f[k] = f[k+1]
				f[k+1] = tmp
			}
		}
	}
	fmt.Println(f,"第一次")
	return f
}

//5. 冒泡排序按value排序，降序
func bubbleValue(f [][]rune) [][]rune{
	for j:=0;j<len(f)-1;j++{
		for k:=0;k<len(f)-j-1;k++{
			if f[k][1] < f[k+1][1] {
				tmp := f[k]
				f[k] = f[k+1]
				f[k+1] = tmp
			}
		}
	}
	fmt.Println(f,"第二次")
	return f
}

//获取top n的字符
var n int
func getTopChr(r [][]rune)  {
	fmt.Println("请输入你需要获取前几位的字符出现次数：")
	fmt.Scan(&n)
	fmt.Printf("前%d位的字符出现的次数如下：\n",n)
	for i,v := range newRune {
		if i < n {
			fmt.Printf("%s出现的次数：%d次\n",string(v[0]),v[1])
		}else {
			break
		}
	}
}

func main() {
	runeString = stringToRune(dreams)
	getMapFromRuneString = elementCount(runeString)
	mapToSlice(getMapFromRuneString)
	bubbleKey(newRune)
	bubbleValue(newRune)
	getTopChr(newRune)
}
