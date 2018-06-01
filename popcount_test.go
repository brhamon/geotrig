package main

import (
	"fmt"
	"testing"
)

var (
	TestData = []TestEntry{
		TestEntry{569162312, 14},
		TestEntry{473027565, 17},
		TestEntry{228507831, 16},
		TestEntry{774109555, 18},
		TestEntry{482436268, 13},
		TestEntry{603686472, 15},
		TestEntry{141340401, 15},
		TestEntry{421325141, 15},
		TestEntry{306313539, 13},
		TestEntry{69327040, 9},
		TestEntry{978593734, 15},
		TestEntry{467745160, 14},
		TestEntry{93749762, 9},
		TestEntry{69873942, 10},
		TestEntry{528936119, 19},
		TestEntry{986169230, 20},
		TestEntry{881742878, 14},
		TestEntry{240038843, 18},
		TestEntry{285036464, 16},
		TestEntry{896180926, 17},
		TestEntry{980527130, 14},
		TestEntry{144235875, 14},
		TestEntry{94867962, 16},
		TestEntry{746605859, 11},
		TestEntry{705090952, 12},
		TestEntry{49631140, 15},
		TestEntry{449850164, 13},
		TestEntry{183023220, 15},
		TestEntry{522214998, 15},
		TestEntry{618783057, 13},
		TestEntry{598534920, 14},
		TestEntry{86083329, 9},
		TestEntry{389145006, 16},
		TestEntry{69555835, 14},
		TestEntry{937890723, 18},
		TestEntry{648777914, 15},
		TestEntry{548875133, 17},
		TestEntry{576765666, 14},
		TestEntry{836266014, 14},
		TestEntry{254027737, 15},
		TestEntry{29940838, 13},
		TestEntry{929842772, 15},
		TestEntry{817961807, 14},
		TestEntry{749232257, 10},
		TestEntry{325879853, 13},
		TestEntry{339988860, 14},
		TestEntry{439949158, 15},
		TestEntry{188607767, 18},
		TestEntry{360377973, 19},
		TestEntry{209953916, 13},
		TestEntry{980254351, 20},
		TestEntry{763328596, 18},
		TestEntry{598672786, 14},
		TestEntry{315656987, 13},
		TestEntry{233952655, 18},
		TestEntry{943556821, 15},
		TestEntry{758036574, 17},
		TestEntry{466922087, 17},
		TestEntry{947981242, 13},
		TestEntry{113894274, 13},
		TestEntry{991712444, 16},
		TestEntry{335379895, 22},
		TestEntry{896126370, 16},
		TestEntry{563204598, 15},
		TestEntry{735519231, 20},
		TestEntry{424216863, 13},
		TestEntry{83387635, 15},
		TestEntry{491067789, 14},
		TestEntry{129934936, 16},
		TestEntry{74496968, 13},
		TestEntry{114552172, 17},
		TestEntry{898992215, 15},
		TestEntry{726578888, 16},
		TestEntry{661887787, 18},
		TestEntry{550631749, 14},
		TestEntry{827581841, 15},
		TestEntry{100624785, 17},
		TestEntry{842400575, 15},
		TestEntry{922516567, 20},
		TestEntry{802204628, 17},
		TestEntry{649344551, 15},
		TestEntry{881836810, 14},
		TestEntry{881236538, 14},
		TestEntry{533791224, 15},
		TestEntry{721851858, 14},
		TestEntry{922003113, 17},
		TestEntry{37331740, 12},
		TestEntry{976286298, 15},
		TestEntry{211668250, 15},
		TestEntry{41355739, 15},
		TestEntry{690811941, 14},
		TestEntry{542139037, 11},
		TestEntry{464393980, 18},
		TestEntry{694932706, 16},
		TestEntry{115567362, 13},
		TestEntry{258964155, 21},
		TestEntry{132085234, 20},
		TestEntry{404884928, 9},
		TestEntry{570943531, 13},
		TestEntry{436853331, 14},
		TestEntry{880023579, 13},
		TestEntry{89839401, 15},
		TestEntry{330150782, 19},
		TestEntry{593203510, 15},
		TestEntry{77946113, 11},
		TestEntry{428013032, 15},
		TestEntry{961883939, 16},
		TestEntry{661037175, 16},
		TestEntry{514584162, 18},
		TestEntry{968458473, 15},
		TestEntry{396907829, 15},
		TestEntry{319359881, 11},
		TestEntry{298552881, 14},
		TestEntry{492633443, 19},
		TestEntry{344449368, 13},
		TestEntry{880684654, 18},
		TestEntry{537547461, 10},
		TestEntry{363872377, 12},
		TestEntry{601865331, 20},
		TestEntry{881142919, 13},
		TestEntry{686860807, 13},
		TestEntry{374060498, 16},
		TestEntry{525541773, 15},
		TestEntry{602234523, 18},
		TestEntry{81687370, 15},
		TestEntry{7473763, 10},
		TestEntry{643589201, 13},
		TestEntry{156055239, 15},
		TestEntry{156473097, 13},
		TestEntry{545910393, 15},
		TestEntry{992035723, 14},
		TestEntry{194016202, 14},
		TestEntry{428030648, 14},
		TestEntry{166509170, 16},
		TestEntry{859999529, 13},
		TestEntry{160735443, 12},
		TestEntry{54029829, 12},
		TestEntry{31011862, 12},
		TestEntry{236484364, 13},
		TestEntry{464856769, 15},
		TestEntry{668514282, 19},
		TestEntry{547879575, 17},
		TestEntry{332739607, 15},
		TestEntry{333888162, 16},
		TestEntry{739785981, 16},
		TestEntry{673285543, 12},
		TestEntry{334628139, 14},
		TestEntry{221883703, 17},
		TestEntry{670773917, 19},
		TestEntry{83638512, 14},
		TestEntry{142625067, 9},
		TestEntry{425634485, 17},
		TestEntry{402585354, 20},
		TestEntry{193663965, 16},
		TestEntry{547083755, 17},
		TestEntry{105189726, 12},
		TestEntry{898449098, 17},
		TestEntry{141341770, 12},
		TestEntry{625200645, 13},
		TestEntry{549729861, 11},
		TestEntry{754120525, 19},
		TestEntry{717955717, 13},
		TestEntry{844896504, 14},
		TestEntry{447518068, 15},
		TestEntry{447587384, 14},
		TestEntry{999672726, 18},
		TestEntry{635639118, 15},
		TestEntry{618802248, 11},
		TestEntry{307333502, 14},
		TestEntry{223460419, 14},
		TestEntry{55849842, 13},
		TestEntry{912236682, 15},
		TestEntry{714828796, 19},
		TestEntry{706858137, 13},
		TestEntry{978625010, 15},
		TestEntry{820242492, 15},
		TestEntry{330704551, 16},
		TestEntry{389443481, 18},
		TestEntry{851979684, 13},
		TestEntry{808646227, 15},
		TestEntry{479864682, 15},
		TestEntry{690990320, 16},
		TestEntry{604809233, 9},
		TestEntry{596442008, 16},
		TestEntry{994291230, 17},
		TestEntry{691577529, 14},
		TestEntry{473273754, 15},
		TestEntry{631852707, 15},
		TestEntry{891752760, 13},
		TestEntry{921368558, 22},
		TestEntry{992749005, 16},
		TestEntry{848216851, 14},
		TestEntry{815321569, 13},
		TestEntry{900059415, 16},
		TestEntry{88189606, 12},
		TestEntry{939487476, 21},
		TestEntry{476463403, 13},
		TestEntry{245235674, 21},
		TestEntry{577564680, 11},
		TestEntry{911258973, 16},
		TestEntry{143966014, 16},
		TestEntry{955674978, 16},
		TestEntry{97595456, 9},
		TestEntry{949597043, 17},
		TestEntry{564101697, 12},
		TestEntry{334325265, 15},
		TestEntry{871490785, 17},
		TestEntry{284040519, 15},
		TestEntry{470751092, 15},
		TestEntry{245975789, 16},
		TestEntry{214105836, 17},
		TestEntry{898561937, 18},
		TestEntry{846196206, 21},
		TestEntry{295441548, 11},
		TestEntry{907699068, 16},
		TestEntry{601815057, 16},
		TestEntry{662449889, 17},
		TestEntry{16723398, 16},
		TestEntry{553231411, 14},
		TestEntry{29906441, 10},
		TestEntry{144796984, 12},
		TestEntry{811308574, 15},
		TestEntry{345650847, 16},
		TestEntry{492278296, 14},
		TestEntry{876218497, 10},
		TestEntry{407516400, 12},
		TestEntry{238635060, 12},
		TestEntry{637221767, 19},
		TestEntry{95899477, 17},
		TestEntry{871392913, 15},
		TestEntry{613120963, 16},
		TestEntry{669020618, 15},
		TestEntry{775082663, 16},
		TestEntry{671337828, 12},
		TestEntry{954091032, 14},
		TestEntry{814663866, 14},
		TestEntry{672833899, 13},
		TestEntry{745090235, 16},
		TestEntry{698275348, 15},
		TestEntry{140911627, 10},
		TestEntry{285925968, 10},
		TestEntry{20808777, 11},
		TestEntry{916584434, 19},
		TestEntry{497273489, 15},
		TestEntry{990490027, 17},
		TestEntry{393850070, 18},
		TestEntry{176642031, 18},
		TestEntry{815832829, 16},
		TestEntry{451047053, 16},
		TestEntry{882398789, 13},
		TestEntry{767988151, 17},
		TestEntry{998085663, 19},
		TestEntry{333244653, 18},
		TestEntry{871065478, 18},
		TestEntry{648966324, 15},
		TestEntry{590458728, 15},
		TestEntry{810860234, 15},
		TestEntry{531893672, 15},
		TestEntry{45163109, 11},
		TestEntry{645032456, 12},
		TestEntry{10354925, 11},
		TestEntry{648879780, 15},
		TestEntry{454163577, 17},
		TestEntry{907986174, 18},
		TestEntry{120258586, 10},
		TestEntry{605674697, 14},
		TestEntry{258113257, 18},
		TestEntry{227721949, 15},
		TestEntry{620545828, 15},
		TestEntry{985967791, 17},
		TestEntry{194085288, 11},
		TestEntry{475823874, 15},
		TestEntry{420945873, 16},
		TestEntry{589080210, 13},
		TestEntry{753458076, 17},
		TestEntry{66692761, 16},
		TestEntry{380123026, 14},
		TestEntry{385443583, 21},
		TestEntry{100997045, 13},
		TestEntry{579430091, 14},
		TestEntry{504747933, 18},
		TestEntry{873419972, 13},
		TestEntry{830589812, 14},
		TestEntry{919422796, 16},
		TestEntry{591073072, 14},
		TestEntry{793572316, 19},
		TestEntry{760912667, 17},
		TestEntry{63028096, 12},
		TestEntry{60493450, 13},
		TestEntry{10095274, 10},
		TestEntry{690228697, 13},
		TestEntry{278912616, 16},
		TestEntry{856290233, 17},
		TestEntry{268485456, 7},
		TestEntry{221781068, 12},
		TestEntry{842094254, 15},
		TestEntry{922436534, 19},
		TestEntry{427793676, 17},
		TestEntry{306802980, 11},
		TestEntry{675146704, 16},
		TestEntry{839350475, 15},
		TestEntry{739270587, 16},
		TestEntry{354200303, 17},
		TestEntry{87631408, 11},
		TestEntry{324445001, 14},
		TestEntry{870829533, 21},
		TestEntry{591251350, 17},
		TestEntry{603011384, 15},
		TestEntry{174934347, 14},
		TestEntry{66695268, 14},
		TestEntry{698677846, 16},
		TestEntry{689395064, 15},
		TestEntry{113397664, 12},
		TestEntry{316392667, 17},
		TestEntry{525886307, 16},
		TestEntry{369041119, 22},
		TestEntry{86712974, 12},
		TestEntry{117731024, 12},
		TestEntry{442502085, 12},
		TestEntry{183526979, 13},
		TestEntry{975172604, 21},
		TestEntry{206380416, 11},
		TestEntry{428586584, 15},
		TestEntry{862597760, 12},
		TestEntry{555360387, 9},
		TestEntry{496051486, 14},
		TestEntry{254403474, 15},
		TestEntry{625967819, 15},
		TestEntry{953177707, 15},
		TestEntry{281764421, 12},
		TestEntry{176629178, 14},
		TestEntry{348866869, 14},
		TestEntry{482899424, 14},
		TestEntry{927839688, 17},
		TestEntry{68179563, 11},
		TestEntry{430796000, 15},
		TestEntry{120519553, 16},
		TestEntry{266417651, 18},
		TestEntry{956974463, 16},
		TestEntry{603634265, 18},
		TestEntry{410093404, 14},
		TestEntry{25558704, 14},
		TestEntry{582718941, 18},
		TestEntry{803045725, 18},
		TestEntry{131097897, 13},
		TestEntry{862576389, 17},
		TestEntry{719007770, 14},
		TestEntry{726237866, 13},
		TestEntry{727287228, 16},
		TestEntry{484387764, 18},
		TestEntry{693550138, 16},
		TestEntry{404697110, 12},
		TestEntry{532557623, 21},
		TestEntry{250868283, 19},
		TestEntry{608834415, 14},
		TestEntry{617393574, 15},
		TestEntry{661995048, 16},
		TestEntry{226205217, 16},
		TestEntry{856735228, 17},
		TestEntry{535078100, 16},
		TestEntry{589175471, 16},
		TestEntry{764630302, 16},
		TestEntry{66928290, 17},
		TestEntry{726988049, 15},
		TestEntry{985416119, 17},
		TestEntry{864589309, 17},
		TestEntry{947879932, 23},
		TestEntry{401339238, 20},
		TestEntry{752171590, 15},
		TestEntry{301726024, 17},
		TestEntry{701252229, 12},
		TestEntry{965898402, 13},
		TestEntry{741457238, 16},
		TestEntry{295189789, 14},
		TestEntry{398926394, 15},
		TestEntry{767468562, 15},
		TestEntry{432265476, 13},
		TestEntry{172200068, 9},
		TestEntry{862242485, 16},
		TestEntry{322835517, 15},
		TestEntry{154282033, 10},
		TestEntry{199679336, 17},
		TestEntry{359947515, 18},
		TestEntry{460619565, 19},
		TestEntry{289180937, 12},
		TestEntry{623635069, 18},
		TestEntry{96504828, 14},
		TestEntry{448876711, 13},
		TestEntry{480595904, 14},
		TestEntry{738379605, 13},
		TestEntry{424530226, 14},
		TestEntry{856633052, 17},
		TestEntry{988383554, 14},
		TestEntry{584305246, 16},
		TestEntry{683695889, 12},
		TestEntry{10178911, 14},
		TestEntry{266294013, 21},
		TestEntry{46666165, 11},
		TestEntry{585069271, 19},
		TestEntry{598597305, 19},
		TestEntry{141361694, 11},
		TestEntry{689055116, 11},
		TestEntry{277334560, 10},
		TestEntry{558143195, 14},
		TestEntry{788580348, 16},
		TestEntry{326501987, 14},
		TestEntry{627309454, 18},
		TestEntry{916712453, 15},
		TestEntry{747757142, 16},
		TestEntry{719070195, 17},
		TestEntry{278761610, 11},
		TestEntry{887050462, 18},
		TestEntry{36151121, 14},
		TestEntry{637184872, 17},
		TestEntry{627028308, 16},
		TestEntry{71593395, 13},
		TestEntry{936855342, 19},
		TestEntry{447214013, 19},
		TestEntry{931361213, 18},
		TestEntry{185708553, 10},
		TestEntry{413718119, 15},
		TestEntry{771690105, 20},
		TestEntry{3608152, 11},
		TestEntry{774529706, 14},
		TestEntry{865116854, 14},
		TestEntry{761696061, 18},
		TestEntry{856469351, 17},
		TestEntry{47075792, 12},
		TestEntry{29067010, 12},
		TestEntry{674589197, 13},
		TestEntry{839367698, 13},
		TestEntry{361732592, 16},
		TestEntry{813140025, 14},
		TestEntry{590838807, 17},
		TestEntry{567218770, 13},
		TestEntry{975751860, 14},
		TestEntry{467218475, 15},
		TestEntry{932961590, 20},
		TestEntry{173543735, 12},
		TestEntry{764492726, 17},
		TestEntry{991067210, 15},
		TestEntry{918648960, 12},
		TestEntry{808684628, 12},
		TestEntry{180747258, 19},
		TestEntry{711548158, 18},
		TestEntry{795387135, 19},
		TestEntry{992071416, 17},
		TestEntry{424180932, 13},
		TestEntry{896526138, 20},
		TestEntry{721928911, 16},
		TestEntry{18227422, 11},
		TestEntry{572742194, 12},
		TestEntry{96493667, 10},
		TestEntry{469051586, 15},
		TestEntry{321965271, 15},
		TestEntry{455003159, 16},
		TestEntry{227508308, 12},
		TestEntry{424632351, 15},
		TestEntry{195338820, 11},
		TestEntry{460687069, 18},
		TestEntry{268601924, 7},
		TestEntry{963395967, 17},
		TestEntry{106239996, 16},
		TestEntry{159348958, 19},
		TestEntry{976324122, 13},
		TestEntry{618385830, 16},
		TestEntry{181197313, 12},
		TestEntry{211389868, 14},
		TestEntry{14648272, 13},
		TestEntry{206684800, 9},
		TestEntry{970756985, 18},
		TestEntry{982670516, 15},
		TestEntry{186515852, 11},
		TestEntry{330198104, 15},
		TestEntry{890394369, 12},
		TestEntry{433208043, 17},
		TestEntry{86004365, 10},
		TestEntry{974444474, 17},
		TestEntry{900332026, 20},
		TestEntry{627119271, 13},
		TestEntry{222761500, 12},
		TestEntry{936860545, 18},
		TestEntry{817587221, 14},
		TestEntry{475595566, 13},
		TestEntry{818312650, 15},
		TestEntry{5641159, 12},
		TestEntry{199459989, 14},
		TestEntry{128233124, 14},
		TestEntry{836955704, 16},
		TestEntry{831688266, 13},
		TestEntry{344554854, 15},
		TestEntry{445295788, 14},
		TestEntry{466126720, 12},
		TestEntry{679511065, 8},
		TestEntry{226360877, 20},
		TestEntry{177103993, 14},
		TestEntry{848720351, 18},
		TestEntry{1642659, 8},
		TestEntry{855485789, 20},
		TestEntry{50677633, 10},
		TestEntry{965366888, 13},
		TestEntry{990432244, 16},
		TestEntry{313002604, 11},
		TestEntry{161228312, 11},
		TestEntry{160631473, 12},
		TestEntry{376003186, 15},
		TestEntry{677023215, 17},
		TestEntry{244582488, 10},
		TestEntry{523911913, 16},
		TestEntry{692950168, 13},
		TestEntry{776783626, 14},
		TestEntry{665435059, 19},
		TestEntry{665739472, 14},
		TestEntry{688443625, 12},
		TestEntry{860527521, 15},
		TestEntry{350664012, 15},
		TestEntry{503943620, 12},
		TestEntry{734291548, 15},
		TestEntry{611672161, 12},
		TestEntry{685160166, 17},
		TestEntry{75833644, 9},
		TestEntry{402404023, 19},
		TestEntry{68684245, 10},
		TestEntry{253151692, 15},
		TestEntry{947071625, 14},
		TestEntry{157657910, 15},
		TestEntry{309840134, 14},
		TestEntry{242175452, 18},
		TestEntry{496560671, 17},
		TestEntry{220289082, 12},
		TestEntry{472830756, 14},
		TestEntry{573199401, 11},
		TestEntry{629829797, 13},
		TestEntry{773637755, 17},
		TestEntry{817379713, 12},
		TestEntry{970950898, 18},
		TestEntry{194507263, 21},
		TestEntry{647917864, 14},
		TestEntry{878284365, 15},
		TestEntry{467744285, 16},
		TestEntry{247902329, 15},
		TestEntry{183492361, 18},
		TestEntry{664845743, 18},
		TestEntry{261399344, 13},
		TestEntry{851004310, 16},
		TestEntry{490306929, 18},
		TestEntry{231286571, 15},
		TestEntry{232058056, 15},
		TestEntry{683622004, 18},
		TestEntry{364988400, 14},
		TestEntry{383116060, 16},
		TestEntry{2898940, 14},
		TestEntry{113526791, 10},
		TestEntry{170735122, 12},
		TestEntry{458228333, 12},
		TestEntry{184836255, 12},
		TestEntry{781720348, 14},
		TestEntry{763616381, 18},
		TestEntry{276795050, 16},
		TestEntry{488153467, 15},
		TestEntry{88062491, 17},
		TestEntry{152641929, 10},
		TestEntry{559406877, 18},
		TestEntry{847837343, 16},
		TestEntry{752227822, 17},
		TestEntry{815042165, 14},
		TestEntry{995076041, 19},
		TestEntry{865731402, 13},
		TestEntry{503565032, 14},
		TestEntry{25975799, 16},
		TestEntry{136551393, 13},
		TestEntry{161912764, 16},
		TestEntry{624735279, 16},
		TestEntry{812959485, 17},
		TestEntry{735702249, 19},
		TestEntry{25689871, 17},
		TestEntry{401603506, 22},
		TestEntry{14805202, 12},
		TestEntry{642672598, 17},
		TestEntry{660380855, 18},
		TestEntry{108554762, 12},
		TestEntry{783104100, 15},
		TestEntry{500421546, 18},
		TestEntry{589389562, 16},
		TestEntry{963214729, 17},
		TestEntry{69745902, 13},
		TestEntry{655905321, 12},
		TestEntry{518400954, 18},
		TestEntry{863926584, 18},
		TestEntry{758323939, 16},
		TestEntry{436052937, 18},
		TestEntry{84894805, 13},
		TestEntry{372970744, 15},
		TestEntry{565416976, 11},
		TestEntry{146698323, 14},
		TestEntry{997024493, 19},
		TestEntry{473894858, 17},
		TestEntry{325587888, 12},
		TestEntry{155111919, 18},
		TestEntry{475799137, 14},
		TestEntry{664653864, 14},
		TestEntry{902626156, 19},
		TestEntry{398328560, 15},
		TestEntry{733671789, 19},
		TestEntry{21799979, 11},
		TestEntry{606397383, 14},
		TestEntry{300969701, 16},
		TestEntry{501848443, 19},
		TestEntry{712877082, 15},
		TestEntry{79972187, 13},
		TestEntry{266170000, 17},
		TestEntry{860870694, 16},
		TestEntry{541303334, 10},
		TestEntry{144886731, 14},
		TestEntry{219694341, 10},
		TestEntry{70739473, 12},
		TestEntry{995445824, 13},
		TestEntry{342813039, 18},
		TestEntry{306160558, 17},
		TestEntry{197817891, 15},
		TestEntry{428702310, 16},
		TestEntry{737536731, 21},
		TestEntry{350485582, 17},
		TestEntry{84179689, 13},
		TestEntry{235297856, 10},
		TestEntry{737680108, 17},
		TestEntry{344595931, 15},
		TestEntry{365211914, 12},
		TestEntry{791561617, 15},
		TestEntry{909884916, 20},
		TestEntry{330421489, 16},
		TestEntry{612009848, 15},
		TestEntry{832599241, 13},
		TestEntry{241752198, 13},
		TestEntry{374801319, 15},
		TestEntry{81821775, 10},
		TestEntry{176064796, 14},
		TestEntry{920398176, 14},
		TestEntry{474020197, 14},
		TestEntry{652836639, 20},
		TestEntry{101357279, 15},
		TestEntry{414670041, 17},
		TestEntry{98354336, 12},
		TestEntry{892344595, 12},
		TestEntry{772678024, 11},
		TestEntry{703889256, 18},
		TestEntry{591407074, 12},
		TestEntry{8042665, 13},
		TestEntry{466294237, 18},
		TestEntry{300567851, 15},
		TestEntry{752155479, 19},
		TestEntry{674576123, 17},
		TestEntry{846610151, 19},
		TestEntry{20380493, 16},
		TestEntry{35119917, 14},
		TestEntry{416395765, 16},
		TestEntry{183751133, 18},
		TestEntry{234410811, 18},
		TestEntry{609225710, 13},
		TestEntry{578666068, 17},
		TestEntry{798625762, 17},
		TestEntry{297168442, 16},
		TestEntry{72614382, 11},
		TestEntry{991068394, 13},
		TestEntry{581535177, 13},
		TestEntry{193472114, 12},
		TestEntry{628711622, 14},
		TestEntry{176691415, 12},
		TestEntry{489230168, 14},
		TestEntry{431289269, 18},
		TestEntry{518584531, 18},
		TestEntry{998991364, 14},
		TestEntry{171772778, 14},
		TestEntry{622348756, 12},
		TestEntry{708507683, 15},
		TestEntry{276139618, 13},
		TestEntry{864458212, 15},
		TestEntry{932745732, 12},
		TestEntry{113345182, 12},
		TestEntry{185569691, 15},
		TestEntry{759313913, 16},
		TestEntry{647705610, 13},
		TestEntry{350469954, 16},
		TestEntry{743002269, 13},
		TestEntry{194424904, 12},
		TestEntry{61956571, 15},
		TestEntry{717448746, 14},
		TestEntry{481299806, 14},
		TestEntry{920893171, 20},
		TestEntry{144502418, 14},
		TestEntry{533188372, 18},
		TestEntry{208125847, 19},
		TestEntry{932254384, 13},
		TestEntry{520350185, 17},
		TestEntry{283392865, 13},
		TestEntry{205328581, 12},
		TestEntry{442656914, 12},
		TestEntry{610517374, 15},
		TestEntry{420978546, 17},
		TestEntry{879283370, 14},
		TestEntry{712614064, 13},
		TestEntry{978467461, 14},
		TestEntry{3312837, 10},
		TestEntry{945815518, 23},
		TestEntry{572601428, 10},
		TestEntry{304061162, 16},
		TestEntry{394212546, 17},
		TestEntry{90550870, 14},
		TestEntry{225649632, 14},
		TestEntry{739916592, 13},
		TestEntry{105283227, 16},
		TestEntry{659707328, 13},
		TestEntry{933470990, 16},
		TestEntry{60593997, 15},
		TestEntry{885435069, 17},
		TestEntry{709468384, 11},
		TestEntry{670046025, 15},
		TestEntry{410786396, 14},
		TestEntry{738718818, 14},
		TestEntry{935592210, 11},
		TestEntry{998876810, 15},
		TestEntry{196452600, 15},
		TestEntry{840160307, 14},
		TestEntry{460608595, 15},
		TestEntry{711618293, 18},
		TestEntry{240884566, 17},
		TestEntry{732431194, 13},
		TestEntry{486701072, 10},
		TestEntry{151714251, 15},
		TestEntry{837086606, 17},
		TestEntry{615569324, 15},
		TestEntry{660544545, 15},
		TestEntry{493648636, 19},
		TestEntry{581634146, 11},
		TestEntry{600624734, 17},
		TestEntry{430071427, 14},
		TestEntry{449574659, 17},
		TestEntry{700877807, 18},
		TestEntry{481994772, 13},
		TestEntry{969658067, 19},
		TestEntry{839556486, 13},
		TestEntry{637461451, 20},
		TestEntry{272858255, 14},
		TestEntry{53508201, 12},
		TestEntry{991233453, 15},
		TestEntry{636370840, 18},
		TestEntry{150823268, 14},
		TestEntry{264004340, 17},
		TestEntry{323401516, 15},
		TestEntry{76038748, 9},
		TestEntry{528060689, 15},
		TestEntry{804347154, 17},
		TestEntry{236847341, 13},
		TestEntry{343756524, 17},
		TestEntry{348324742, 11},
		TestEntry{108786928, 17},
		TestEntry{116004287, 17},
		TestEntry{568769500, 18},
		TestEntry{726165366, 16},
		TestEntry{719686065, 15},
		TestEntry{513619662, 18},
		TestEntry{337577299, 13},
		TestEntry{969877069, 17},
		TestEntry{125998172, 12},
		TestEntry{214312110, 13},
		TestEntry{537882238, 16},
		TestEntry{408776796, 14},
		TestEntry{866363055, 18},
		TestEntry{168259416, 14},
		TestEntry{324965246, 18},
		TestEntry{975346617, 16},
		TestEntry{650526683, 19},
		TestEntry{82403420, 12},
		TestEntry{352242739, 17},
		TestEntry{345479221, 14},
		TestEntry{553147713, 13},
		TestEntry{908419389, 15},
		TestEntry{468480334, 17},
		TestEntry{787283186, 20},
		TestEntry{671449894, 10},
		TestEntry{714803387, 15},
		TestEntry{671901172, 13},
		TestEntry{598036106, 14},
		TestEntry{157594186, 12},
		TestEntry{692663760, 13},
		TestEntry{127125810, 14},
		TestEntry{749845075, 16},
		TestEntry{285784955, 15},
		TestEntry{34897237, 13},
		TestEntry{375388631, 21},
		TestEntry{167882145, 11},
		TestEntry{537204044, 9},
		TestEntry{296034167, 17},
		TestEntry{669515744, 21},
		TestEntry{490081740, 15},
		TestEntry{931477642, 14},
		TestEntry{16383957, 19},
		TestEntry{383530736, 16},
		TestEntry{640219559, 16},
		TestEntry{479441093, 14},
		TestEntry{864341073, 13},
		TestEntry{238033572, 11},
		TestEntry{142676757, 9},
		TestEntry{17310145, 7},
		TestEntry{479545833, 15},
		TestEntry{883058643, 16},
		TestEntry{481570199, 16},
		TestEntry{301059912, 15},
		TestEntry{45511693, 13},
		TestEntry{353836321, 13},
		TestEntry{492810604, 18},
		TestEntry{682849989, 16},
		TestEntry{369898502, 10},
		TestEntry{145301096, 12},
		TestEntry{948473062, 13},
		TestEntry{880172156, 17},
		TestEntry{760162793, 17},
		TestEntry{811663621, 8},
		TestEntry{637405350, 15},
		TestEntry{657770096, 14},
		TestEntry{397357480, 16},
		TestEntry{237753270, 17},
		TestEntry{739132980, 12},
		TestEntry{385562479, 21},
		TestEntry{522010567, 20},
		TestEntry{857318643, 16},
		TestEntry{130796880, 15},
		TestEntry{256435114, 15},
		TestEntry{115621273, 15},
		TestEntry{746325763, 13},
		TestEntry{683137451, 18},
		TestEntry{939264259, 15},
		TestEntry{684205006, 13},
		TestEntry{260871786, 15},
		TestEntry{176942116, 13},
		TestEntry{631517178, 17},
		TestEntry{20541844, 12},
		TestEntry{85822589, 15},
		TestEntry{451526350, 19},
		TestEntry{640594226, 14},
		TestEntry{588612029, 14},
		TestEntry{603361697, 15},
		TestEntry{281150159, 11},
		TestEntry{327534167, 15},
		TestEntry{167435798, 17},
		TestEntry{767055104, 12},
		TestEntry{794826624, 12},
		TestEntry{908101986, 11},
		TestEntry{223894779, 17},
		TestEntry{231075008, 14},
		TestEntry{687140948, 14},
		TestEntry{353214600, 10},
		TestEntry{941203961, 15},
		TestEntry{775498110, 17},
		TestEntry{706038888, 12},
		TestEntry{892642794, 16},
		TestEntry{904612771, 18},
		TestEntry{903916793, 17},
		TestEntry{499127023, 15},
		TestEntry{670856551, 20},
		TestEntry{759075729, 16},
		TestEntry{955073235, 16},
		TestEntry{65971396, 14},
		TestEntry{556992541, 11},
		TestEntry{915213357, 14},
		TestEntry{323190625, 16},
		TestEntry{109936961, 10},
		TestEntry{729252080, 15},
		TestEntry{208216081, 9},
		TestEntry{495380094, 18},
		TestEntry{509635169, 13},
		TestEntry{969771854, 17},
		TestEntry{750973518, 15},
		TestEntry{714493518, 15},
		TestEntry{551641829, 13},
		TestEntry{801830140, 20},
		TestEntry{975529608, 13},
		TestEntry{891874822, 14},
		TestEntry{575626060, 15},
		TestEntry{932966840, 19},
		TestEntry{392579286, 15},
		TestEntry{450937853, 17},
		TestEntry{234062607, 16},
		TestEntry{882282926, 15},
		TestEntry{419883356, 14},
		TestEntry{452439519, 22},
		TestEntry{401312356, 16},
		TestEntry{821322844, 14},
		TestEntry{538940771, 14},
		TestEntry{789605978, 15},
		TestEntry{393536762, 18},
		TestEntry{958648430, 15},
		TestEntry{276266789, 17},
		TestEntry{679893245, 15},
		TestEntry{754605338, 17},
		TestEntry{25997444, 9},
		TestEntry{440498854, 14},
		TestEntry{846926157, 15},
		TestEntry{959751587, 15},
		TestEntry{266105505, 16},
		TestEntry{683136316, 17},
		TestEntry{823839678, 17},
		TestEntry{269563276, 10},
		TestEntry{688179810, 11},
		TestEntry{858548876, 14},
		TestEntry{322646031, 14},
		TestEntry{223841312, 12},
		TestEntry{798704682, 15},
		TestEntry{528363673, 17},
		TestEntry{452138775, 17},
		TestEntry{904688485, 18},
		TestEntry{84058947, 10},
		TestEntry{12087746, 11},
		TestEntry{176951933, 13},
		TestEntry{117634526, 16},
		TestEntry{777496078, 16},
		TestEntry{266291813, 18},
		TestEntry{106432276, 10},
		TestEntry{246436030, 14},
		TestEntry{179293186, 13},
		TestEntry{875731345, 13},
		TestEntry{152383544, 10},
		TestEntry{728092090, 18},
		TestEntry{624333980, 14},
		TestEntry{778850229, 18},
		TestEntry{977421680, 12},
		TestEntry{676878631, 13},
		TestEntry{478641526, 18},
		TestEntry{213779764, 13},
		TestEntry{217006046, 21},
		TestEntry{255062080, 13},
		TestEntry{890677941, 16},
		TestEntry{54836024, 13},
		TestEntry{358038667, 16},
		TestEntry{601359338, 14},
		TestEntry{903833683, 18},
		TestEntry{77905291, 14},
		TestEntry{814907104, 14},
		TestEntry{371105542, 15},
		TestEntry{730454082, 13},
		TestEntry{429828265, 15},
		TestEntry{415101019, 17},
		TestEntry{114107331, 14},
		TestEntry{848292353, 14},
		TestEntry{976635665, 14},
		TestEntry{402520959, 25},
		TestEntry{73542726, 10},
		TestEntry{743405717, 16},
		TestEntry{677366035, 15},
		TestEntry{168322917, 12},
		TestEntry{143485697, 11},
		TestEntry{762599891, 17},
		TestEntry{941493602, 13},
		TestEntry{768174132, 14},
		TestEntry{628380072, 13},
		TestEntry{721374256, 16},
		TestEntry{915520702, 18},
		TestEntry{373663454, 16},
		TestEntry{162199274, 17},
		TestEntry{912782417, 17},
		TestEntry{115239523, 16},
		TestEntry{236759311, 14},
		TestEntry{43799811, 11},
		TestEntry{615020150, 15},
		TestEntry{734048718, 15},
		TestEntry{317752699, 15},
		TestEntry{956292935, 22},
		TestEntry{804659618, 16},
		TestEntry{763931933, 15},
		TestEntry{133581684, 17},
		TestEntry{914479850, 16},
		TestEntry{481320366, 16},
		TestEntry{901062785, 12},
		TestEntry{8397438, 9},
		TestEntry{21154942, 13},
		TestEntry{199287060, 12},
		TestEntry{869871372, 16},
		TestEntry{196543367, 15},
		TestEntry{628245788, 13},
		TestEntry{313493072, 13},
		TestEntry{604362876, 13},
		TestEntry{487469442, 12},
		TestEntry{724667740, 15},
		TestEntry{181622833, 13},
		TestEntry{426721277, 22},
		TestEntry{494560886, 18},
		TestEntry{196965909, 17},
		TestEntry{63877826, 14},
		TestEntry{643310715, 14},
		TestEntry{915248924, 16},
		TestEntry{370837473, 14},
		TestEntry{335232141, 18},
		TestEntry{938448562, 20},
		TestEntry{300828525, 17},
		TestEntry{101266966, 11},
		TestEntry{502600964, 14},
		TestEntry{104173550, 17},
		TestEntry{409863577, 13},
		TestEntry{131809743, 17},
		TestEntry{693216641, 12},
	}
)

func TestPopCount(t *testing.T) {
	for _, entry := range TestData {
		guess := PopCount(int64(entry.Value))
		if guess != entry.BinaryWeight {
			t.Errorf("Incorrect result. For Value %d, expected %d. Got %d\n",
				entry.Value, entry.BinaryWeight, guess)
		}
	}
	fmt.Printf("Tested %d binary weights.\n", len(TestData))
}
