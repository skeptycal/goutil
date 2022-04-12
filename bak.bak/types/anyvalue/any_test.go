package anyvalue

import (
	"fmt"
	"reflect"
	"testing"
)

var (
	// TRun         = types.TRun
	// BRun         = types.BRun
	// TName        = types.TName

	// LimitResult  = benchmark.LimitResult
	globalReturn bool
	global       Any
)

/* IsComparable benchmarks

/// Using the raw struct's method is roughly twice as fast as using an external function.

invalid.IsComparable()-8         	35839928	        33.66 ns/op	       0 B/op	       0 allocs/op
IsComparable(invalid)-8          	19477659	        60.77 ns/op	       0 B/op	       0 allocs/op
bool.IsComparable()-8            	35884406	        33.35 ns/op	       0 B/op	       0 allocs/op
IsComparable(bool)-8             	19895932	        60.32 ns/op	       0 B/op	       0 allocs/op
int.IsComparable()-8             	36053312	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(int)-8              	19916199	        60.32 ns/op	       0 B/op	       0 allocs/op
uint.IsComparable()-8            	36006660	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(uint)-8             	19909026	        63.69 ns/op	       0 B/op	       0 allocs/op
int.IsComparable()#01-8          	36054621	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(int)#01-8           	19890436	        60.35 ns/op	       0 B/op	       0 allocs/op
Int8.IsComparable()-8            	36000630	        33.32 ns/op	       0 B/op	       0 allocs/op
IsComparable(Int8)-8             	19913238	        60.38 ns/op	       0 B/op	       0 allocs/op
Int16.IsComparable()-8           	36064645	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Int16)-8            	19915455	        60.32 ns/op	       0 B/op	       0 allocs/op
Int32.IsComparable()-8           	36063470	        33.32 ns/op	       0 B/op	       0 allocs/op
IsComparable(Int32)-8            	19898778	        60.30 ns/op	       0 B/op	       0 allocs/op
Int64.IsComparable()-8           	36053312	        33.33 ns/op	       0 B/op	       0 allocs/op
IsComparable(Int64)-8            	19905985	        60.30 ns/op	       0 B/op	       0 allocs/op
Uint.IsComparable()-8            	36052410	        33.30 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uint)-8             	19886083	        60.23 ns/op	       0 B/op	       0 allocs/op
Uint8.IsComparable()-8           	36066180	        33.27 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uint8)-8            	19933057	        60.34 ns/op	       0 B/op	       0 allocs/op
Uint16.IsComparable()-8          	36053990	        33.33 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uint16)-8           	19907911	        60.41 ns/op	       0 B/op	       0 allocs/op
Uint32.IsComparable()-8          	36045372	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uint32)-8           	19915483	        60.32 ns/op	       0 B/op	       0 allocs/op
Uint64.IsComparable()-8          	36063741	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uint64)-8           	19655535	        60.34 ns/op	       0 B/op	       0 allocs/op
Uintptr.IsComparable()-8         	36010622	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Uintptr)-8          	19908530	        60.32 ns/op	       0 B/op	       0 allocs/op
Float32.IsComparable()-8         	36003194	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Float32)-8          	19896495	        60.31 ns/op	       0 B/op	       0 allocs/op
Float64.IsComparable()-8         	36073363	        33.33 ns/op	       0 B/op	       0 allocs/op
IsComparable(Float64)-8          	19902326	        60.29 ns/op	       0 B/op	       0 allocs/op
Complex64.IsComparable()-8       	35676193	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Complex64)-8        	19902216	        60.31 ns/op	       0 B/op	       0 allocs/op
Complex128.IsComparable()-8      	36019989	        33.40 ns/op	       0 B/op	       0 allocs/op
IsComparable(Complex128)-8       	19905792	        60.34 ns/op	       0 B/op	       0 allocs/op
Array.IsComparable()-8           	36068845	        33.32 ns/op	       0 B/op	       0 allocs/op
IsComparable(Array)-8            	19890999	        60.36 ns/op	       0 B/op	       0 allocs/op
Chan.IsComparable()-8            	35942760	        33.33 ns/op	       0 B/op	       0 allocs/op
IsComparable(Chan)-8             	19917190	        60.30 ns/op	       0 B/op	       0 allocs/op
Func.IsComparable()-8            	36072052	        33.32 ns/op	       0 B/op	       0 allocs/op
IsComparable(Func)-8             	19896316	        60.31 ns/op	       0 B/op	       0 allocs/op
Map.IsComparable()-8             	35987493	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Map)-8              	19907264	        60.29 ns/op	       0 B/op	       0 allocs/op
Ptr.IsComparable()-8             	36040815	        33.33 ns/op	       0 B/op	       0 allocs/op
IsComparable(Ptr)-8              	19899480	        60.36 ns/op	       0 B/op	       0 allocs/op
Slice.IsComparable()-8           	36040815	        33.32 ns/op	       0 B/op	       0 allocs/op
IsComparable(Slice)-8            	19921765	        60.52 ns/op	       0 B/op	       0 allocs/op
String.IsComparable()-8          	36040861	        33.34 ns/op	       0 B/op	       0 allocs/op
IsComparable(String)-8           	19869632	        60.30 ns/op	       0 B/op	       0 allocs/op
Struct.IsComparable()-8          	36057285	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(Struct)-8           	19908489	        60.32 ns/op	       0 B/op	       0 allocs/op
UnsafePointer.IsComparable()-8   	36012063	        33.31 ns/op	       0 B/op	       0 allocs/op
IsComparable(UnsafePointer)-8    	19909204	        60.31 ns/op	       0 B/op	       0 allocs/op

/// Raw Struct method is faster than global function which is faster than the interface method
struct_method-8         	66181033	        18.17 ns/op	       0 B/op	       0 allocs/op
interface_method-8      	56819302	        21.20 ns/op	       0 B/op	       0 allocs/op
global_function-8       	78668318	        15.27 ns/op	       0 B/op	       0 allocs/op
struct_method#01-8      	65958074	        24.09 ns/op	       0 B/op	       0 allocs/op
interface_method#01-8   	43465137	        25.53 ns/op	       0 B/op	       0 allocs/op
global_function#01-8    	60631065	        21.37 ns/op	       0 B/op	       0 allocs/op
struct_method#02-8      	57729892	        24.57 ns/op	       0 B/op	       0 allocs/op
interface_method#02-8   	49745655	        25.56 ns/op	       0 B/op	       0 allocs/op
global_function#02-8    	61368256	        20.38 ns/op	       0 B/op	       0 allocs/op
struct_method#03-8      	58688912	        20.44 ns/op	       0 B/op	       0 allocs/op
interface_method#03-8   	52602644	        23.66 ns/op	       0 B/op	       0 allocs/op
global_function#03-8    	60365084	        18.59 ns/op	       0 B/op	       0 allocs/op
struct_method#04-8      	66399070	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#04-8   	56994830	        21.70 ns/op	       0 B/op	       0 allocs/op
global_function#04-8    	69400767	        17.44 ns/op	       0 B/op	       0 allocs/op
struct_method#05-8      	66414688	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#05-8   	56725632	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#05-8    	62037544	        19.18 ns/op	       0 B/op	       0 allocs/op
struct_method#06-8      	66384226	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#06-8   	56818294	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#06-8    	61668914	        19.07 ns/op	       0 B/op	       0 allocs/op
struct_method#07-8      	66522688	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#07-8   	56925096	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#07-8    	69185862	        17.46 ns/op	       0 B/op	       0 allocs/op
struct_method#08-8      	66521151	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#08-8   	56958981	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#08-8    	67236932	        17.62 ns/op	       0 B/op	       0 allocs/op
struct_method#09-8      	66533599	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#09-8   	56955940	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#09-8    	68495918	        17.49 ns/op	       0 B/op	       0 allocs/op
struct_method#10-8      	66498417	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#10-8   	55875350	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#10-8    	62110460	        19.71 ns/op	       0 B/op	       0 allocs/op
struct_method#11-8      	66540822	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#11-8   	56933312	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#11-8    	67269603	        17.97 ns/op	       0 B/op	       0 allocs/op
struct_method#12-8      	66302920	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#12-8   	57012204	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#12-8    	67829500	        18.02 ns/op	       0 B/op	       0 allocs/op
struct_method#13-8      	66403663	        18.08 ns/op	       0 B/op	       0 allocs/op
interface_method#13-8   	56877652	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#13-8    	66270423	        18.02 ns/op	       0 B/op	       0 allocs/op
struct_method#14-8      	66538054	        18.04 ns/op	       0 B/op	       0 allocs/op
interface_method#14-8   	57017508	        21.06 ns/op	       0 B/op	       0 allocs/op
global_function#14-8    	57525447	        20.18 ns/op	       0 B/op	       0 allocs/op
struct_method#15-8      	66444873	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#15-8   	57003740	        21.09 ns/op	       0 B/op	       0 allocs/op
global_function#15-8    	65435361	        18.50 ns/op	       0 B/op	       0 allocs/op
struct_method#16-8      	66526066	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#16-8   	56988174	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#16-8    	60615118	        19.93 ns/op	       0 B/op	       0 allocs/op
struct_method#17-8      	66390805	        18.06 ns/op	       0 B/op	       0 allocs/op
interface_method#17-8   	57021799	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#17-8    	69717787	        17.44 ns/op	       0 B/op	       0 allocs/op
struct_method#18-8      	64895119	        18.06 ns/op	       0 B/op	       0 allocs/op
interface_method#18-8   	56931846	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#18-8    	68608998	        17.42 ns/op	       0 B/op	       0 allocs/op
struct_method#19-8      	66259600	        18.06 ns/op	       0 B/op	       0 allocs/op
interface_method#19-8   	56883381	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#19-8    	64560675	        18.76 ns/op	       0 B/op	       0 allocs/op
struct_method#20-8      	66394018	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#20-8   	56983324	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#20-8    	66526221	        17.99 ns/op	       0 B/op	       0 allocs/op
struct_method#21-8      	66404122	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#21-8   	56754249	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#21-8    	66667899	        17.96 ns/op	       0 B/op	       0 allocs/op
struct_method#22-8      	66472632	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#22-8   	56882596	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#22-8    	65461983	        18.06 ns/op	       0 B/op	       0 allocs/op
struct_method#23-8      	66541438	        18.06 ns/op	       0 B/op	       0 allocs/op
interface_method#23-8   	56917782	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#23-8    	66575280	        17.96 ns/op	       0 B/op	       0 allocs/op
struct_method#24-8      	66499954	        18.07 ns/op	       0 B/op	       0 allocs/op
interface_method#24-8   	57005208	        21.13 ns/op	       0 B/op	       0 allocs/op
global_function#24-8    	57827726	        20.82 ns/op	       0 B/op	       0 allocs/op
struct_method#25-8      	66501336	        18.09 ns/op	       0 B/op	       0 allocs/op
interface_method#25-8   	56912947	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#25-8    	67786077	        17.95 ns/op	       0 B/op	       0 allocs/op
struct_method#26-8      	66529604	        18.09 ns/op	       0 B/op	       0 allocs/op
interface_method#26-8   	57010624	        21.08 ns/op	       0 B/op	       0 allocs/op
global_function#26-8    	65645065	        18.66 ns/op	       0 B/op	       0 allocs/op
struct_method#27-8      	66478465	        18.08 ns/op	       0 B/op	       0 allocs/op
interface_method#27-8   	56716915	        21.07 ns/op	       0 B/op	       0 allocs/op
global_function#27-8    	56448120	        21.19 ns/op	       0 B/op	       0 allocs/op
*/

/* Benchmarks

struct_method-8         	70093677	        18.84 ns/op	       0 B/op	       0 allocs/op
interface_method-8      	58739425	        20.64 ns/op	       0 B/op	       0 allocs/op
global_function-8       	79790548	        15.08 ns/op	       0 B/op	       0 allocs/op

struct_method#01-8      	70213418	        16.95 ns/op	       0 B/op	       0 allocs/op
interface_method#01-8   	58615930	        20.39 ns/op	       0 B/op	       0 allocs/op
global_function#01-8    	65146137	        17.75 ns/op	       0 B/op	       0 allocs/op

struct_method#02-8      	71826717	        16.97 ns/op	       0 B/op	       0 allocs/op
interface_method#02-8   	58519220	        20.40 ns/op	       0 B/op	       0 allocs/op
global_function#02-8    	65026696	        18.47 ns/op	       0 B/op	       0 allocs/op

struct_method#03-8      	70534989	        16.98 ns/op	       0 B/op	       0 allocs/op
interface_method#03-8   	59613507	        20.50 ns/op	       0 B/op	       0 allocs/op
global_function#03-8    	60967864	        19.76 ns/op	       0 B/op	       0 allocs/op

struct_method#04-8      	69642934	        17.00 ns/op	       0 B/op	       0 allocs/op
interface_method#04-8   	57491798	        20.49 ns/op	       0 B/op	       0 allocs/op
global_function#04-8    	64666206	        18.53 ns/op	       0 B/op	       0 allocs/op

struct_method#05-8      	70551403	        17.03 ns/op	       0 B/op	       0 allocs/op
interface_method#05-8   	58570749	        20.49 ns/op	       0 B/op	       0 allocs/op
global_function#05-8    	62887312	        19.13 ns/op	       0 B/op	       0 allocs/op
struct_method#06-8      	71051814	        16.94 ns/op	       0 B/op	       0 allocs/op
interface_method#06-8   	59895679	        20.47 ns/op	       0 B/op	       0 allocs/op
global_function#06-8    	66613160	        18.01 ns/op	       0 B/op	       0 allocs/op
struct_method#07-8      	70922337	        17.04 ns/op	       0 B/op	       0 allocs/op
interface_method#07-8   	57870138	        20.57 ns/op	       0 B/op	       0 allocs/op
global_function#07-8    	64483339	        18.59 ns/op	       0 B/op	       0 allocs/op
struct_method#08-8      	71028860	        17.10 ns/op	       0 B/op	       0 allocs/op
interface_method#08-8   	58949713	        20.35 ns/op	       0 B/op	       0 allocs/op
global_function#08-8    	62821330	        19.13 ns/op	       0 B/op	       0 allocs/op
struct_method#09-8      	71832810	        16.91 ns/op	       0 B/op	       0 allocs/op
interface_method#09-8   	60082237	        20.35 ns/op	       0 B/op	       0 allocs/op
global_function#09-8    	60973028	        19.70 ns/op	       0 B/op	       0 allocs/op
struct_method#10-8      	71775163	        17.03 ns/op	       0 B/op	       0 allocs/op
interface_method#10-8   	57753626	        20.36 ns/op	       0 B/op	       0 allocs/op
global_function#10-8    	64338705	        18.65 ns/op	       0 B/op	       0 allocs/op
struct_method#11-8      	71812387	        16.94 ns/op	       0 B/op	       0 allocs/op
interface_method#11-8   	60081109	        20.48 ns/op	       0 B/op	       0 allocs/op
global_function#11-8    	61975336	        19.26 ns/op	       0 B/op	       0 allocs/op
struct_method#12-8      	70796982	        17.09 ns/op	       0 B/op	       0 allocs/op
interface_method#12-8   	59537222	        20.60 ns/op	       0 B/op	       0 allocs/op
global_function#12-8    	67110028	        17.90 ns/op	       0 B/op	       0 allocs/op
struct_method#13-8      	70173045	        17.07 ns/op	       0 B/op	       0 allocs/op
interface_method#13-8   	58300657	        20.55 ns/op	       0 B/op	       0 allocs/op
global_function#13-8    	58837950	        20.31 ns/op	       0 B/op	       0 allocs/op
struct_method#14-8      	71054090	        17.08 ns/op	       0 B/op	       0 allocs/op
interface_method#14-8   	57976615	        20.50 ns/op	       0 B/op	       0 allocs/op
global_function#14-8    	62079802	        19.38 ns/op	       0 B/op	       0 allocs/op
struct_method#15-8      	68954122	        17.08 ns/op	       0 B/op	       0 allocs/op
interface_method#15-8   	58665003	        20.59 ns/op	       0 B/op	       0 allocs/op
global_function#15-8    	65163531	        18.48 ns/op	       0 B/op	       0 allocs/op
struct_method#16-8      	70669122	        17.09 ns/op	       0 B/op	       0 allocs/op
interface_method#16-8   	58312344	        20.58 ns/op	       0 B/op	       0 allocs/op
global_function#16-8    	60625833	        19.82 ns/op	       0 B/op	       0 allocs/op
struct_method#17-8      	70022902	        17.11 ns/op	       0 B/op	       0 allocs/op
interface_method#17-8   	57134468	        20.53 ns/op	       0 B/op	       0 allocs/op
global_function#17-8    	66810494	        17.91 ns/op	       0 B/op	       0 allocs/op
struct_method#18-8      	70188265	        17.06 ns/op	       0 B/op	       0 allocs/op
interface_method#18-8   	58651742	        20.53 ns/op	       0 B/op	       0 allocs/op
global_function#18-8    	67417782	        17.90 ns/op	       0 B/op	       0 allocs/op
struct_method#19-8      	70054755	        17.09 ns/op	       0 B/op	       0 allocs/op
interface_method#19-8   	57838524	        20.76 ns/op	       0 B/op	       0 allocs/op
global_function#19-8    	60725797	        19.70 ns/op	       0 B/op	       0 allocs/op
struct_method#20-8      	70273381	        16.92 ns/op	       0 B/op	       0 allocs/op
interface_method#20-8   	59852985	        20.30 ns/op	       0 B/op	       0 allocs/op
global_function#20-8    	65092102	        18.51 ns/op	       0 B/op	       0 allocs/op
struct_method#21-8      	70812819	        16.97 ns/op	       0 B/op	       0 allocs/op
interface_method#21-8   	57330091	        20.43 ns/op	       0 B/op	       0 allocs/op
global_function#21-8    	59440884	        20.22 ns/op	       0 B/op	       0 allocs/op
struct_method#22-8      	69935163	        16.97 ns/op	       0 B/op	       0 allocs/op
interface_method#22-8   	59814201	        20.38 ns/op	       0 B/op	       0 allocs/op
global_function#22-8    	65134790	        18.44 ns/op	       0 B/op	       0 allocs/op
struct_method#23-8      	70039422	        16.95 ns/op	       0 B/op	       0 allocs/op
interface_method#23-8   	59842042	        20.38 ns/op	       0 B/op	       0 allocs/op
global_function#23-8    	67180003	        17.84 ns/op	       0 B/op	       0 allocs/op
struct_method#24-8      	71363434	        17.01 ns/op	       0 B/op	       0 allocs/op
interface_method#24-8   	59187771	        20.40 ns/op	       0 B/op	       0 allocs/op
global_function#24-8    	62910114	        19.31 ns/op	       0 B/op	       0 allocs/op
struct_method#25-8      	70927573	        16.97 ns/op	       0 B/op	       0 allocs/op
interface_method#25-8   	59417826	        20.36 ns/op	       0 B/op	       0 allocs/op
global_function#25-8    	64898629	        18.48 ns/op	       0 B/op	       0 allocs/op
struct_method#26-8      	70664095	        16.92 ns/op	       0 B/op	       0 allocs/op
interface_method#26-8   	59854352	        20.35 ns/op	       0 B/op	       0 allocs/op
global_function#26-8    	64786364	        18.49 ns/op	       0 B/op	       0 allocs/op
struct_method#27-8      	70198698	        16.92 ns/op	       0 B/op	       0 allocs/op
interface_method#27-8   	59796568	        20.35 ns/op	       0 B/op	       0 allocs/op
global_function#27-8    	59468376	        20.21 ns/op	       0 B/op	       0 allocs/op
*/

func Test_any(t *testing.T) {

	for _, tt := range reflectTests {
		if tt.want != reflect.Func {
			TRun(t, tt.name, NewAnyValue(tt.a).Interface(), tt.a)
		}
	}
}

func Test_any_TypeOf(t *testing.T) {
	for _, tt := range reflectTests {
		TRun(t, tt.name, NewAnyValue(tt.a).TypeOf(), TypeOf(tt.a))
	}
}
func Test_any_Interface(t *testing.T) {
	for _, tt := range reflectTests {
		if tt.name == "Func" {
			continue
		}

		if tt.name == "Struct" {
			continue
		}

		A := NewAnyValue(tt.a)

		want := tt.a
		got := A.Interface()

		t.Run(tt.name+".Interface()", func(t *testing.T) {
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Interface(%v) = %v(%T), want %v(%T)", tt.name, got, got, want, want)
			}
		})
	}
}
func Test_any_Kind(t *testing.T) {
	for _, tt := range reflectTests {
		A := NewAnyValue(tt.a)
		want := KindOf(tt.a)
		got := A.Kind()
		t.Run(tt.name+".Kind()", func(t *testing.T) {
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Kind(%v) = %v, want %v", tt.a, got, want)
			}
		})
	}
}
func Test_any_Is_all(t *testing.T) {

	for _, tt := range reflectTests {
		// a := new_any(tt.a)

		name := fmt.Sprintf("IsComparable(%v)", tt.want)
		name2 := fmt.Sprintf("IsComparable(%v type %T)", tt.want, tt.a)
		A := NewAnyValue(tt.a)
		want := IsComparable(tt.a)
		got := A.IsComparable()
		t.Run(name, func(t *testing.T) {
			if got != want {
				t.Errorf("%v = %v, want %v", name2, got, want)
			}
		})

		name = fmt.Sprintf("IsOrdered(%v)", tt.want)
		name2 = fmt.Sprintf("IsOrdered(%v type %T)", tt.want, tt.a)
		want = IsOrdered(tt.a)
		got = A.IsOrdered()
		t.Run("IsOrdered"+name, func(t *testing.T) {
			if got != want {
				t.Errorf("%v = %v, want %v", name2, got, want)
			}
		})

		name = fmt.Sprintf("IsDeepComparable(%v)", tt.want)
		name2 = fmt.Sprintf("IsDeepComparable(%v type %T)", tt.want, tt.a)
		want = IsDeepComparable(tt.a)
		got = A.IsDeepComparable()
		t.Run("IsDeepComparable"+name, func(t *testing.T) {
			if got != want {
				t.Errorf("%v = %v, want %v", name2, got, want)
			}
		})

		name = fmt.Sprintf("IsIterable(%v)", tt.want)
		name2 = fmt.Sprintf("IsIterable(%v type %T)", tt.want, tt.a)
		want = IsIterable(tt.a)
		got = A.IsIterable()
		t.Run("IsIterable"+name, func(t *testing.T) {
			if got != want {
				t.Errorf("%v = %v, want %v", name2, got, want)
			}
		})

		name = fmt.Sprintf("HasAlternate(%v)", tt.want)
		name2 = fmt.Sprintf("HasAlternate(%v type %T)", tt.want, tt.a)
		want = HasAlternate(tt.a)
		got = A.HasAlternate()
		t.Run("HasAlternate"+name, func(t *testing.T) {
			if got != want {
				t.Errorf("%v = %v, want %v", name2, got, want)
			}
		})

	}
}
func Test_any_Elem(t *testing.T) {
	LimitResult = true

	for _, tt := range reflectTests {

		v := ValueOf(tt.a)
		a := NewAnyValue(tt.a)

		name := TName(tt.name, tt.name, v)
		TRun(t, name, a.Elem(), Elem(v))

	}

	LimitResult = false
}
func Test_any_Indirect(t *testing.T) {
	LimitResult = true

	for _, tt := range reflectTests {
		want := ValueOf(tt.a)
		got := NewAnyValue(tt.a).Indirect().ValueOf()
		name := TName(tt.name, tt.name, tt.a)
		TRun(t, name, got, want)
	}
	LimitResult = false
}
func Test_any_String(t *testing.T) {

	LimitResult = true

	var extraTests = []struct {
		name string
		a    interface{}
		want reflect.Kind
	}{
		{"pointer int", &LimitResult, reflect.Ptr},
		{"pointer int", &reflectTests, reflect.Slice},
	}

	tests := append(reflectTests, extraTests...)

	for _, tt := range tests {

		// want := Elem(tt.a)

		a := NewAnyValue(tt.a)
		got := a.Elem()

		// test raw struct (*any) as well as interface (AnyValue)
		name := TName(tt.name, a.String(), a.ValueOf())
		TRun(t, name, got, nil)

	}
	LimitResult = false

	TRun(t, "AnyValue.String", NewAnyValue("fake").String(), "fake")
}

/*

1. Test object returns a list of values to loop on
- This list should be an interface so that it can be replaced
  with AI generated tests, random tests, premade tests, etc.

2. common function runs tests on that data
*/
