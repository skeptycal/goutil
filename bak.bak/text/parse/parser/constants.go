package parser

const (
	defaultNL  = "\n"
	defaultSep = "\t"
)

var nlbyte []byte = []byte{'\n'}

type Cases byte

const (
	none         Cases = iota // Now is the time for ALL good men to come to the aid of their country.
	upper                     // NOW IS THE TIME FOR ALL GOOD MEN TO COME TO THE AID OF THEIR COUNTRY.
	lower                     // now is the time for all good men to come to the aid of their country.
	title                     // Now Is The Time For All Good Men To Come To The Aid Of Their Country.
	reverse                   // nOW IS THE TIME FOR all GOOD MEN TO COME TO THE AID OF THEIR COUNTRY.
	camel                     // nowIsTheTimeForALLGoodMenToComeToTheAidOfTheirCountry.
	snake                     // now_is_the_time_for_all_good_men_to_come_to_the_aid_of_their_country.
	snakeAllCaps              // NOW_IS_THE_TIME_FOR_ALL_GOOD_MEN_TO_COME_TO_THE_AID_OF_THEIR_COUNTRY.
	Pascal                    // NowIsTheTimeForALLGoodMenToComeToTheAidOfTheirCountry.
	kehab                     // now-is-the-time-for-all-good-men-to-come-to-the-aid-of-their-country.
	snakeCamel                // now_Is_The_Time_For_All_Good_Men_To_Come_To_The_Aid_Of_Their_Country.
	snakePascal               // Now_Is_The_Time_For_All_Good_Men_To_Come_To_The_Aid_Of_Their_Country.
	kehabCamel                // now-Is-The-Time-For-All-Good-Men-To-Come-To-The-Aid-Of-Their-Country
	kehabPascal               // Now-Is-The-Time-For-All-Good-Men-To-Come-To-The-Aid-Of-Their-Country
)
