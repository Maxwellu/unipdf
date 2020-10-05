//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package arithmetic ;import (_g "bytes";_gf "github.com/unidoc/unipdf/v3/common";_a "github.com/unidoc/unipdf/v3/internal/jbig2/bitmap";_gd "github.com/unidoc/unipdf/v3/internal/jbig2/errors";_fa "io";);func (_fag *Encoder )codeLPS (_aaa *codingContext ,_adce uint32 ,_cea uint16 ,_edg byte ){_fag ._fe -=_cea ;if _fag ._fe < _cea {_fag ._ee +=uint32 (_cea );}else {_fag ._fe =_cea ;};if _cbaf [_edg ]._aga ==1{_aaa .flipMps (_adce );};_aaa ._ga [_adce ]=_cbaf [_edg ]._eaf ;_fag .renormalize ();};func (_geaf *Encoder )emit (){if _geaf ._cc ==_cfbb {_geaf ._adf =append (_geaf ._adf ,_geaf ._add );_geaf ._add =make ([]byte ,_cfbb );_geaf ._cc =0;};_geaf ._add [_geaf ._cc ]=_geaf ._bg ;_geaf ._cc ++;};func (_ed *codingContext )mps (_af uint32 )int {return int (_ed ._cf [_af ])};func (_afd *Encoder )Refine (iTemp ,iTarget *_a .Bitmap ,ox ,oy int )error {for _eb :=0;_eb < iTarget .Height ;_eb ++{var _cfb int ;_cbf :=_eb +oy ;var (_aeb ,_bga ,_bdc ,_ab ,_fgf uint16 ;_ffe ,_cg ,_dfd ,_fef ,_ffc byte ;);if _cbf >=1&&(_cbf -1)< iTemp .Height {_ffe =iTemp .Data [(_cbf -1)*iTemp .RowStride ];};if _cbf >=0&&_cbf < iTemp .Height {_cg =iTemp .Data [_cbf *iTemp .RowStride ];};if _cbf >=-1&&_cbf +1< iTemp .Height {_dfd =iTemp .Data [(_cbf +1)*iTemp .RowStride ];};if _eb >=1{_fef =iTarget .Data [(_eb -1)*iTarget .RowStride ];};_ffc =iTarget .Data [_eb *iTarget .RowStride ];_dbe :=uint (6+ox );_aeb =uint16 (_ffe >>_dbe );_bga =uint16 (_cg >>_dbe );_bdc =uint16 (_dfd >>_dbe );_ab =uint16 (_fef >>6);_gc :=uint (2-ox );_ffe <<=_gc ;_cg <<=_gc ;_dfd <<=_gc ;_fef <<=2;for _cfb =0;_cfb < iTarget .Width ;_cfb ++{_be :=(_aeb <<10)|(_bga <<7)|(_bdc <<4)|(_ab <<1)|_fgf ;_bb :=_ffc >>7;_cgf :=_afd .encodeBit (_afd ._dg ,uint32 (_be ),_bb );if _cgf !=nil {return _cgf ;};_aeb <<=1;_bga <<=1;_bdc <<=1;_ab <<=1;_aeb |=uint16 (_ffe >>7);_bga |=uint16 (_cg >>7);_bdc |=uint16 (_dfd >>7);_ab |=uint16 (_fef >>7);_fgf =uint16 (_bb );_ge :=_cfb %8;_gae :=_cfb /8+1;if _ge ==5+ox {_ffe ,_cg ,_dfd =0,0,0;if _gae < iTemp .RowStride &&_cbf >=1&&(_cbf -1)< iTemp .Height {_ffe =iTemp .Data [(_cbf -1)*iTemp .RowStride +_gae ];};if _gae < iTemp .RowStride &&_cbf >=0&&_cbf < iTemp .Height {_cg =iTemp .Data [_cbf *iTemp .RowStride +_gae ];};if _gae < iTemp .RowStride &&_cbf >=-1&&(_cbf +1)< iTemp .Height {_dfd =iTemp .Data [(_cbf +1)*iTemp .RowStride +_gae ];};}else {_ffe <<=1;_cg <<=1;_dfd <<=1;};if _ge ==5&&_eb >=1{_fef =0;if _gae < iTarget .RowStride {_fef =iTarget .Data [(_eb -1)*iTarget .RowStride +_gae ];};}else {_fef <<=1;};if _ge ==7{_ffc =0;if _gae < iTarget .RowStride {_ffc =iTarget .Data [_eb *iTarget .RowStride +_gae ];};}else {_ffc <<=1;};_aeb &=7;_bga &=7;_bdc &=7;_ab &=7;};};return nil ;};func (_faa *Encoder )Flush (){_faa ._cc =0;_faa ._adf =nil ;_faa ._ae =-1};func (_gea *Encoder )byteOut (){if _gea ._bg ==0xff{_gea .rBlock ();return ;};if _gea ._ee < 0x8000000{_gea .lBlock ();return ;};_gea ._bg ++;if _gea ._bg !=0xff{_gea .lBlock ();return ;};_gea ._ee &=0x7ffffff;_gea .rBlock ();};func (_agc *Encoder )Final (){_agc .flush ()};func (_gcd *Encoder )encodeInteger (_fea Class ,_dga int )error {const _cfc ="E\u006e\u0063\u006f\u0064er\u002ee\u006e\u0063\u006f\u0064\u0065I\u006e\u0074\u0065\u0067\u0065\u0072";if _dga > 2000000000||_dga < -2000000000{return _gd .Errorf (_cfc ,"\u0061\u0072\u0069\u0074\u0068\u006d\u0065\u0074i\u0063\u0020\u0065nc\u006f\u0064\u0065\u0072\u0020\u002d \u0069\u006e\u0076\u0061\u006c\u0069\u0064\u0020\u0069\u006e\u0074\u0065\u0067\u0065\u0072 \u0076\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",_dga );};_ecf :=_gcd ._fac [_fea ];_de :=uint32 (1);var _fff int ;for ;;_fff ++{if _ag [_fff ]._ad <=_dga &&_ag [_fff ]._c >=_dga {break ;};};if _dga < 0{_dga =-_dga ;};_dga -=int (_ag [_fff ]._ffb );_dbae :=_ag [_fff ]._fd ;for _fbb :=uint8 (0);_fbb < _ag [_fff ]._ff ;_fbb ++{_dcf :=_dbae &1;if _aba :=_gcd .encodeBit (_ecf ,_de ,_dcf );_aba !=nil {return _gd .Wrap (_aba ,_cfc ,"");};_dbae >>=1;if _de &0x100> 0{_de =(((_de <<1)|uint32 (_dcf ))&0x1ff)|0x100;}else {_de =(_de <<1)|uint32 (_dcf );};};_dga <<=32-_ag [_fff ]._e ;for _feb :=uint8 (0);_feb < _ag [_fff ]._e ;_feb ++{_daf :=uint8 ((uint32 (_dga )&0x80000000)>>31);if _daa :=_gcd .encodeBit (_ecf ,_de ,_daf );_daa !=nil {return _gd .Wrap (_daa ,_cfc ,"\u006d\u006f\u0076\u0065 \u0064\u0061\u0074\u0061\u0020\u0074\u006f\u0020\u0074\u0068e\u0020t\u006f\u0070\u0020\u006f\u0066\u0020\u0077o\u0072\u0064");};_dga <<=1;if _de &0x100!=0{_de =(((_de <<1)|uint32 (_daf ))&0x1ff)|0x100;}else {_de =(_de <<1)|uint32 (_daf );};};return nil ;};func (_cgc *Encoder )renormalize (){for {_cgc ._fe <<=1;_cgc ._ee <<=1;_cgc ._ac --;if _cgc ._ac ==0{_cgc .byteOut ();};if (_cgc ._fe &0x8000)!=0{break ;};};};type Class int ;func (_faf *Encoder )EncodeInteger (proc Class ,value int )(_ede error ){_gf .Log .Trace ("\u0045\u006eco\u0064\u0065\u0020I\u006e\u0074\u0065\u0067er:\u0027%d\u0027\u0020\u0077\u0069\u0074\u0068\u0020Cl\u0061\u0073\u0073\u003a\u0020\u0027\u0025s\u0027",value ,proc );if _ede =_faf .encodeInteger (proc ,value );_ede !=nil {return _gd .Wrap (_ede ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u006e\u0074\u0065\u0067\u0065\u0072","");};return nil ;};func (_dfb *Encoder )flush (){_dfb .setBits ();_dfb ._ee <<=_dfb ._ac ;_dfb .byteOut ();_dfb ._ee <<=_dfb ._ac ;_dfb .byteOut ();_dfb .emit ();if _dfb ._bg !=0xff{_dfb ._ae ++;_dfb ._bg =0xff;_dfb .emit ();};_dfb ._ae ++;_dfb ._bg =0xac;_dfb ._ae ++;_dfb .emit ();};func (_cad *Encoder )encodeIAID (_ccg ,_gfd int )error {if _cad ._ccb ==nil {_cad ._ccb =_dc (1<<uint (_ccg ));};_eba :=uint32 (1<<uint32 (_ccg +1))-1;_gfd <<=uint (32-_ccg );_agg :=uint32 (1);for _gg :=0;_gg < _ccg ;_gg ++{_bde :=_agg &_eba ;_eab :=uint8 ((uint32 (_gfd )&0x80000000)>>31);if _aae :=_cad .encodeBit (_cad ._ccb ,_bde ,_eab );_aae !=nil {return _aae ;};_agg =(_agg <<1)|uint32 (_eab );_gfd <<=1;};return nil ;};func (_bgf *Encoder )WriteTo (w _fa .Writer )(int64 ,error ){const _dcg ="\u0045n\u0063o\u0064\u0065\u0072\u002e\u0057\u0072\u0069\u0074\u0065\u0054\u006f";var _fge int64 ;for _eeb ,_faaa :=range _bgf ._adf {_adc ,_gfe :=w .Write (_faaa );if _gfe !=nil {return 0,_gd .Wrapf (_gfe ,_dcg ,"\u0066\u0061\u0069\u006c\u0065\u0064\u0020\u0061\u0074\u0020\u0069'\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u0020\u0063h\u0075\u006e\u006b",_eeb );};_fge +=int64 (_adc );};_bgf ._add =_bgf ._add [:_bgf ._cc ];_ecg ,_gde :=w .Write (_bgf ._add );if _gde !=nil {return 0,_gd .Wrap (_gde ,_dcg ,"\u0062u\u0066f\u0065\u0072\u0065\u0064\u0020\u0063\u0068\u0075\u006e\u006b\u0073");};_fge +=int64 (_ecg );return _fge ,nil ;};func (_db *Encoder )Init (){_db ._dg =_dc (_eg );_db ._fe =0x8000;_db ._ee =0;_db ._ac =12;_db ._ae =-1;_db ._bg =0;_db ._cc =0;_db ._add =make ([]byte ,_cfbb );for _dbc :=0;_dbc < len (_db ._fac );_dbc ++{_db ._fac [_dbc ]=_dc (512);};_db ._ccb =nil ;};func (_adg *codingContext )flipMps (_gdc uint32 ){_adg ._cf [_gdc ]=1-_adg ._cf [_gdc ]};func (_cba *Encoder )rBlock (){if _cba ._ae >=0{_cba .emit ();};_cba ._ae ++;_cba ._bg =uint8 (_cba ._ee >>20);_cba ._ee &=0xfffff;_cba ._ac =7;};func (_ebe *Encoder )setBits (){_ced :=_ebe ._ee +uint32 (_ebe ._fe );_ebe ._ee |=0xffff;if _ebe ._ee >=_ced {_ebe ._ee -=0x8000;};};func (_cdf *Encoder )lBlock (){if _cdf ._ae >=0{_cdf .emit ();};_cdf ._ae ++;_cdf ._bg =uint8 (_cdf ._ee >>19);_cdf ._ee &=0x7ffff;_cdf ._ac =8;};type state struct{_adfc uint16 ;_gfc ,_eaf uint8 ;_aga uint8 ;};func (_gfb *Encoder )codeMPS (_abc *codingContext ,_adb uint32 ,_eded uint16 ,_gff byte ){_gfb ._fe -=_eded ;if _gfb ._fe &0x8000!=0{_gfb ._ee +=uint32 (_eded );return ;};if _gfb ._fe < _eded {_gfb ._fe =_eded ;}else {_gfb ._ee +=uint32 (_eded );};_abc ._ga [_adb ]=_cbaf [_gff ]._gfc ;_gfb .renormalize ();};type codingContext struct{_ga []byte ;_cf []byte ;};func (_fbd *Encoder )code0 (_fdb *codingContext ,_fgd uint32 ,_abd uint16 ,_cec byte ){if _fdb .mps (_fgd )==0{_fbd .codeMPS (_fdb ,_fgd ,_abd ,_cec );}else {_fbd .codeLPS (_fdb ,_fgd ,_abd ,_cec );};};type intEncRangeS struct{_ad ,_c int ;_fd ,_ff uint8 ;_ffb uint16 ;_e uint8 ;};const _ce =0x9b25;const (_eg =65536;_cfbb =20*1024;);var _cbaf =[]state {{0x5601,1,1,1},{0x3401,2,6,0},{0x1801,3,9,0},{0x0AC1,4,12,0},{0x0521,5,29,0},{0x0221,38,33,0},{0x5601,7,6,1},{0x5401,8,14,0},{0x4801,9,14,0},{0x3801,10,14,0},{0x3001,11,17,0},{0x2401,12,18,0},{0x1C01,13,20,0},{0x1601,29,21,0},{0x5601,15,14,1},{0x5401,16,14,0},{0x5101,17,15,0},{0x4801,18,16,0},{0x3801,19,17,0},{0x3401,20,18,0},{0x3001,21,19,0},{0x2801,22,19,0},{0x2401,23,20,0},{0x2201,24,21,0},{0x1C01,25,22,0},{0x1801,26,23,0},{0x1601,27,24,0},{0x1401,28,25,0},{0x1201,29,26,0},{0x1101,30,27,0},{0x0AC1,31,28,0},{0x09C1,32,29,0},{0x08A1,33,30,0},{0x0521,34,31,0},{0x0441,35,32,0},{0x02A1,36,33,0},{0x0221,37,34,0},{0x0141,38,35,0},{0x0111,39,36,0},{0x0085,40,37,0},{0x0049,41,38,0},{0x0025,42,39,0},{0x0015,43,40,0},{0x0009,44,41,0},{0x0005,45,42,0},{0x0001,45,43,0},{0x5601,46,46,0}};type Encoder struct{_ee uint32 ;_fe uint16 ;_ac ,_bg uint8 ;_ae int ;_aa int ;_adf [][]byte ;_add []byte ;_cc int ;_dg *codingContext ;_fac [13]*codingContext ;_ccb *codingContext ;};func (_edb *Encoder )code1 (_dfg *codingContext ,_cebb uint32 ,_dfe uint16 ,_acg byte ){if _dfg .mps (_cebb )==1{_edb .codeMPS (_dfg ,_cebb ,_dfe ,_acg );}else {_edb .codeLPS (_dfg ,_cebb ,_dfe ,_acg );};};var _ag =[]intEncRangeS {{0,3,0,2,0,2},{-1,-1,9,4,0,0},{-3,-2,5,3,2,1},{4,19,2,3,4,4},{-19,-4,3,3,4,4},{20,83,6,4,20,6},{-83,-20,7,4,20,6},{84,339,14,5,84,8},{-339,-84,15,5,84,8},{340,4435,30,6,340,12},{-4435,-340,31,6,340,12},{4436,2000000000,62,6,4436,32},{-2000000000,-4436,63,6,4436,32}};func _dc (_b int )*codingContext {return &codingContext {_ga :make ([]byte ,_b ),_cf :make ([]byte ,_b )}};const (IAAI Class =iota ;IADH ;IADS ;IADT ;IADW ;IAEX ;IAFS ;IAIT ;IARDH ;IARDW ;IARDX ;IARDY ;IARI ;);func (_age *Encoder )dataSize ()int {return _cfbb *len (_age ._adf )+_age ._cc };var _ _fa .WriterTo =&Encoder {};func New ()*Encoder {_dd :=&Encoder {};_dd .Init ();return _dd };func (_efg *Encoder )encodeOOB (_fbaf Class )error {_abf :=_efg ._fac [_fbaf ];_gga :=_efg .encodeBit (_abf ,1,1);if _gga !=nil {return _gga ;};_gga =_efg .encodeBit (_abf ,3,0);if _gga !=nil {return _gga ;};_gga =_efg .encodeBit (_abf ,6,0);if _gga !=nil {return _gga ;};_gga =_efg .encodeBit (_abf ,12,0);if _gga !=nil {return _gga ;};return nil ;};func (_fb *Encoder )DataSize ()int {return _fb .dataSize ()};func (_afg *Encoder )encodeBit (_fc *codingContext ,_cbfc uint32 ,_gdd uint8 )error {const _agf ="\u0045\u006e\u0063\u006f\u0064\u0065\u0072\u002e\u0065\u006e\u0063\u006fd\u0065\u0042\u0069\u0074";_afg ._aa ++;if _cbfc >=uint32 (len (_fc ._ga )){return _gd .Errorf (_agf ,"\u0061r\u0069\u0074h\u006d\u0065\u0074i\u0063\u0020\u0065\u006e\u0063\u006f\u0064e\u0072\u0020\u002d\u0020\u0069\u006ev\u0061\u006c\u0069\u0064\u0020\u0063\u0074\u0078\u0020\u006e\u0075m\u0062\u0065\u0072\u003a\u0020\u0027\u0025\u0064\u0027",_cbfc );};_efa :=_fc ._ga [_cbfc ];_fee :=_fc .mps (_cbfc );_ebd :=_cbaf [_efa ]._adfc ;_gf .Log .Trace ("\u0045\u0043\u003a\u0020\u0025d\u0009\u0020D\u003a\u0020\u0025d\u0009\u0020\u0049\u003a\u0020\u0025d\u0009\u0020\u004dPS\u003a \u0025\u0064\u0009\u0020\u0051\u0045\u003a \u0025\u0030\u0034\u0058\u0009\u0020\u0020\u0041\u003a\u0020\u0025\u0030\u0034\u0058\u0009\u0020\u0043\u003a %\u0030\u0038\u0058\u0009\u0020\u0043\u0054\u003a\u0020\u0025\u0064\u0009\u0020\u0042\u003a\u0020\u0025\u0030\u0032\u0058\u0009\u0020\u0042\u0050\u003a\u0020\u0025\u0064",_afg ._aa ,_gdd ,_efa ,_fee ,_ebd ,_afg ._fe ,_afg ._ee ,_afg ._ac ,_afg ._bg ,_afg ._ae );if _gdd ==0{_afg .code0 (_fc ,_cbfc ,_ebd ,_efa );}else {_afg .code1 (_fc ,_cbfc ,_ebd ,_efa );};return nil ;};func (_ada *Encoder )EncodeBitmap (bm *_a .Bitmap ,duplicateLineRemoval bool )error {_gf .Log .Trace ("\u0045n\u0063\u006f\u0064\u0065 \u0042\u0069\u0074\u006d\u0061p\u0020[\u0025d\u0078\u0025\u0064\u005d\u002c\u0020\u0025s",bm .Width ,bm .Height ,bm );var (_da ,_eda uint8 ;_aeg ,_ceb ,_gdf uint16 ;_afc ,_bd ,_ca byte ;_cb ,_ea ,_facd int ;_fg ,_cff []byte ;);for _ec :=0;_ec < bm .Height ;_ec ++{_afc ,_bd =0,0;if _ec >=2{_afc =bm .Data [(_ec -2)*bm .RowStride ];};if _ec >=1{_bd =bm .Data [(_ec -1)*bm .RowStride ];if duplicateLineRemoval {_ea =_ec *bm .RowStride ;_fg =bm .Data [_ea :_ea +bm .RowStride ];_facd =(_ec -1)*bm .RowStride ;_cff =bm .Data [_facd :_facd +bm .RowStride ];if _g .Equal (_fg ,_cff ){_eda =_da ^1;_da =1;}else {_eda =_da ;_da =0;};};};if duplicateLineRemoval {if _ffbd :=_ada .encodeBit (_ada ._dg ,_ce ,_eda );_ffbd !=nil {return _ffbd ;};if _da !=0{continue ;};};_ca =bm .Data [_ec *bm .RowStride ];_aeg =uint16 (_afc >>5);_ceb =uint16 (_bd >>4);_afc <<=3;_bd <<=4;_gdf =0;for _cb =0;_cb < bm .Width ;_cb ++{_dba :=uint32 (_aeg <<11|_ceb <<4|_gdf );_df :=(_ca &0x80)>>7;_ef :=_ada .encodeBit (_ada ._dg ,_dba ,_df );if _ef !=nil {return _ef ;};_aeg <<=1;_ceb <<=1;_gdf <<=1;_aeg |=uint16 ((_afc &0x80)>>7);_ceb |=uint16 ((_bd &0x80)>>7);_gdf |=uint16 (_df );_cef :=_cb %8;_fba :=_cb /8+1;if _cef ==4&&_ec >=2{_afc =0;if _fba < bm .RowStride {_afc =bm .Data [(_ec -2)*bm .RowStride +_fba ];};}else {_afc <<=1;};if _cef ==3&&_ec >=1{_bd =0;if _fba < bm .RowStride {_bd =bm .Data [(_ec -1)*bm .RowStride +_fba ];};}else {_bd <<=1;};if _cef ==7{_ca =0;if _fba < bm .RowStride {_ca =bm .Data [_ec *bm .RowStride +_fba ];};}else {_ca <<=1;};_aeg &=31;_ceb &=127;_gdf &=15;};};return nil ;};func (_abg *Encoder )Reset (){_abg ._fe =0x8000;_abg ._ee =0;_abg ._ac =12;_abg ._ae =-1;_abg ._bg =0;_abg ._ccb =nil ;_abg ._dg =_dc (_eg );};func (_d Class )String ()string {switch _d {case IAAI :return "\u0049\u0041\u0041\u0049";case IADH :return "\u0049\u0041\u0044\u0048";case IADS :return "\u0049\u0041\u0044\u0053";case IADT :return "\u0049\u0041\u0044\u0054";case IADW :return "\u0049\u0041\u0044\u0057";case IAEX :return "\u0049\u0041\u0045\u0058";case IAFS :return "\u0049\u0041\u0046\u0053";case IAIT :return "\u0049\u0041\u0049\u0054";case IARDH :return "\u0049\u0041\u0052D\u0048";case IARDW :return "\u0049\u0041\u0052D\u0057";case IARDX :return "\u0049\u0041\u0052D\u0058";case IARDY :return "\u0049\u0041\u0052D\u0059";case IARI :return "\u0049\u0041\u0052\u0049";default:return "\u0055N\u004b\u004e\u004f\u0057\u004e";};};func (_gaa *Encoder )EncodeIAID (symbolCodeLength ,value int )(_cd error ){_gf .Log .Trace ("\u0045\u006e\u0063\u006f\u0064\u0065\u0020\u0049A\u0049\u0044\u002e S\u0079\u006d\u0062\u006f\u006c\u0043o\u0064\u0065\u004c\u0065\u006e\u0067\u0074\u0068\u003a\u0020\u0027\u0025\u0064\u0027\u002c \u0056\u0061\u006c\u0075\u0065\u003a\u0020\u0027%\u0064\u0027",symbolCodeLength ,value );if _cd =_gaa .encodeIAID (symbolCodeLength ,value );_cd !=nil {return _gd .Wrap (_cd ,"\u0045\u006e\u0063\u006f\u0064\u0065\u0049\u0041\u0049\u0044","");};return nil ;};func (_ba *Encoder )EncodeOOB (proc Class )(_bf error ){_gf .Log .Trace ("E\u006e\u0063\u006f\u0064\u0065\u0020O\u004f\u0042\u0020\u0077\u0069\u0074\u0068\u0020\u0043l\u0061\u0073\u0073:\u0020'\u0025\u0073\u0027",proc );if _bf =_ba .encodeOOB (proc );_bf !=nil {return _gd .Wrap (_bf ,"\u0045n\u0063\u006f\u0064\u0065\u004f\u004fB","");};return nil ;};