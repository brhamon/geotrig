package main

import (
	"fmt"
	"math"
)

const (
	// Citation: http://earth-info.nga.mil/GandG/publications/tr8350.2/wgs84fin.pdf
	// see Tables 3.1 and 3.3
	WGS84_b float64 = 6356752.3142
	WGS84_a float64 = 6378137.0
	WGS84_f float64 = 298.257223563
)

var (
	KGTerm = []float64{
		1.0,
		0.25,
		0.015625,
		0.00390625,
		0.00152587890625,
		0.0007476806640625,
		0.00042057037353515625,
		0.0002596378326416015625,
		0.000171401537954807281494140625,
		0.00011902884580194950103759765625,
		0.000085998341091908514499664306640625,
		0.00006414339077309705317020416259765625,
		0.00004910978356065243133343756198883056640625,
		0.0000384305850644750535138882696628570556640625,
		0.00003063662712410319954869919456541538238525390625,
		0.0000248156679705235916344463475979864597320556640625,
		0.00002038083668282259820758728352529942640103399753570556640625,
		0.0000169428927787132498940236846607376719475723803043365478515625,
		0.00001423673629322432803595045724964762712261290289461612701416015625,
		0.0000120775636836563724681712674036138110977844917215406894683837890625,
		0.0000103338654268284836930790406722170671205418557292432524263858795166015625,
		0.000008910322740479661959848764661248389506997824582867906428873538970947265625,
		0.0000077367006853028469806331474150612307651153631837814828031696379184722900390625,
		0.000006760472385219737271829248379228835389744001194145539557212032377719879150390625,
		0.00000594182143232203471156867533330659360426718854954197812645588783198036253452301025390625,
		0.0000052501934175997498711420815245097061087304878023752918725364224883378483355045318603515625,
		0.00000466187662561279565111395626492152528367673861446119666640530709855738678015768527984619140625,
		0.0000041582788419817837752220165449454345894523995666027340635528819490218666032887995243072509765625,
		0.0000037246828020174842552929351003672594903609025454678188726148103937507727323463768698275089263916,
		0.0000033493357538950326611953414621316765631217985136861332014446496554982424243007699260488152503967,
		0.0000030227755178902669767287956695738380982174231586017352143037963140871637879314448582590557634830,
		0.0000027373261128449582065538339557197009417000130112103642770529435404103582585300102891778806224465,
		0.0000024867164223379124723112344114338396494301143590609778991489264926432966503882246792067611806942,
		0.0000022657891368822714882009387922362051351212405626981224246377615356522599645066262056408712410871,
		0.0000020702766226919543766541882346881415864808047961504254420619685311701553525174082436921887961923,
		0.0000018966268896457516728164593847989933840229250469222979202890156604944545668266623685580072461443,
		0.0000017418674038586851300692830113865755210904988712185687497098772298638306698807367933458087382124,
		0.0000016034977324418611652080452265156550770301688841878752862102796047742093511447761459562129016305,
		0.0000014794043310565578513493201198237406346076471578665490651341031879919947424256426734419422702197,
		0.0000013677924658437110312031436019080442257836974462523567540071220303180424763550690397946951462830,
		0.0000012671314890605254225005372524551240960424284623172223741419103933993240378608131776473043003612,
		0.0000011761105923894615053280566615961376388162992316064522363205923208217104878479082453445606987737,
		0.0000010936028339947926497246853396729392075217884436748771432113670942334527367871493760920723844465,
		0.0000010186357386952577831197075858581501082500812044992196646272455262539556386866782114519046317539,
		0.0000009503671503193746749793242907832043559022258138567745450583482602253137254017626649974187712322,
		0.0000008880653037984378907306796983874165147930798993928304359934120964994320478476471125142546517847,
		0.0000008310923052206316791679718680206434562471627933708187480510181021233460835303890333442120861043,
		0.0000007788903779461352348562669804299398439545897568926833468323314739342948073466672233050498285457,
		0.0000007309703644592148053680396173761447168362898011463561487362017055184153416603000015587430519848,
		0.0000006869020761395682651443729224093821396759178941426347607605393994485317012165980335347413623659,
		0.0000006463061634397197806743404826949876552210711465988050463995915209411234776746970897528381478501,
		0.0000006088472422022965753930422021235653603250402064412618473435598324436708193665615317827342067550,
		0.0000005742280619180498673802166701056296450328897139337381753653526119413725987757298618450140202577,
		0.0000005421845415529183911567033333170723481803067795588312835930069295199378693851653706224415931750,
		0.0000005124815304030285718880876414455352056488239235799138290134517659428425077136015270158109194748,
		0.0000004849091769904358776484888766041266586341640579393746635020668816760003199019028002317371253775,
		0.0000004592798096160210987198418640731528086122850105530700236820835954394578922795366047156623713815,
		0.0000004354252488672665402683265317978851765860236699772526748066291150669098715586465456064693811782,
		0.0000004131944859383268766859587904672410686553906244009764718048340643794123179200622577919892633966,
		0.0000003924516716844565458325053866654167719741123964164689629143156062494777294235006721702856943708,
		0.0000003730743703950365038820254331988118188578905968434058078704212731909097665332153264818778382113,
		0.0000003549520397180940561323140392050775441310527238577982830726307208852776944287061433962558496983,
		0.0000003379847043127351115916499640999961188620410334288517600459408418627309263872715039974363875802,
		0.0000003220817958898569855927231233855405191650175607675168983204232172172622943633806112356522491623,
		0.0000003071611365221567016532164796691327277803588493037384970859749004528639739641004669529459468482,
		0.0000002931480456192819787553093846499077968265921822733726757100407792546889370453832207978736790956,
		0.0000002799745538998204435529788492859914857088682567269969407995172525009916552670008136649113804998,
		0.0000002675787101511928398202645373466751997243199016313207006605321658592959342858655025230309757606,
		0.0000002559039686345399082818263084518456751688740668229039724256138344444791188139421969144623124043,
		0.0000002448986467320147988046778235420545804428024505275900492258355457230955650275202971416758897064,
		0.0000002345154439037339672839284729622868581801509792832825323428422121263663602041596151557201415254,
		0.0000002247110142662192016411814137127724849681956492130679333166065453527833984082804961527310481260,
		0.0000002154455861606242258790667286855531333744549431908277190522499386650601246023835138894890995271,
		0.0000002066826229779792078720695972457720033953006724202118608978916774142341193466945240911129009303,
		0.0000001983885202753840780455744741641871973788438932448390419730721565757063714054169270003491938485,
		0.0000001905323348724788685549697249872853843626416750723434159109384991753083990977624166911353657721,
		0.0000001830855421790124377072750547283034460801163360578729300830482003199022579799785064473639307266,
		0.0000001760218184864084412280139367034932903555714529623697368368857318052829897200830631432933456104,
		0.0000001693168453709868179120060093808380355824117415514510671274925253053036450672840411374652337029,
		0.0000001629481337140665879000137948796119934652876979159434340545588816079121964725804794234738919930,
		0.0000001568948651530479423885718761713889073017920494503941291410477294044307707364311030198909360834,
		0.0000001511377490448942627467415638427405489062873343299578562267500246560514523314934733823297803355,
		0.0000001456588932552314167407156482885067581870863322860959842078222556926498250998156723506607018172,
		0.0000001404416872876412945051558302865922506268216273228801061263474202169405284902381549820258450639,
		0.0000001354706964429575624611276743038716703272115505904694901250640772181904013657431182463737823082,
		0.0000001307315658511295314698404743481203118946575409833080833943914203992426333456473987810767617575,
		0.0000001262109333516127145859286704927212083566560988380294135285023444437117648385963817126261963412,
		0.0000001218963503149196521075155322327143894027276055662180631849298802311591926161116659287522330299,
		0.0000001177762096001817622651676253936243853445969301069001941200208672985008870030864556295721068682,
		0.0000001138396799332649434847480913924929554721083507298263617259701761462122732126474783378249518003,
		0.0000001100766460688042411862244739270188827773358802473682125466950508791569230703405200261024047824,
		0.0000001064776541688973762784633006006403822928577146179786528863861890236404713228106690664275797498,
		0.0000001030338618923454319547121984575135740872020199846289770560874863717948216270853121835194334884,
		0.0000000997369927423042019519989829501581709621432664835599437111606495289928541296525615595410541130,
		0.0000000965792942679199103612258146069817621429196835502444282909255667193803879749705443462905323964,
		0.0000000935534997577532228648671886701259069356165765669943881691239928700834567062810239680175520047,
		0.0000000906527931002252298707660819901683898016265118150392399031379706302151464031321738596341952897,
		0.0000000878707765195375866435173088820101240395668184058998435249860853055818566248449578747293835255,
		0.0000000852014409250378895482188733482401892531711895772949622933727272893486718403490690305027802723,
		0.0000000826391386382656297844868548889611569317374370900071661362487999994256516357839340344063927113,
	}
)

func KummerGauss(h float64) (sum float64) {
	var hProd float64 = 1
	sum = 0
	for term, kg := range KGTerm {
		if term != 0 {
			hProd *= h
		}
		sum += hProd * kg
	}
	return
}

func EllipsePerimeter(a float64, b float64) float64 {
	diff := a - b
	sum := a + b
	return KummerGauss((diff*diff)/(sum*sum)) * sum * math.Pi
}

func ellipse() {
	ep := CircumferencePrecise(WGS84_a)
	fmt.Printf(" Equatorial perimeter: %50.42f\n", ep)
	pp := EllipsePerimeterPrecise(WGS84_a, WGS84_b)
	fmt.Printf("      Polar perimeter: %50.42f\n", pp)
	ep /= 4
	pp /= 4
	fmt.Printf("90-deg equatorial arc: %50.42f\n", ep)
	fmt.Printf("     90-deg polar arc: %50.42f\n", pp)
	fmt.Printf("                error: %50.42f\n", ep-pp)
	ef := pp / ep
	fmt.Printf("       scaling factor: %50.42f\n", ef)

}