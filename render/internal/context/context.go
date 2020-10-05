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

package context ;import (_gb "errors";_dgc "github.com/golang/freetype/truetype";_cg "github.com/unidoc/unipdf/v3/core";_e "github.com/unidoc/unipdf/v3/internal/textencoding";_de "github.com/unidoc/unipdf/v3/internal/transform";_c "github.com/unidoc/unipdf/v3/model";_ce "golang.org/x/image/font";_dg "image";_d "image/color";);func NewTextFont (font *_c .PdfFont ,size float64 )(*TextFont ,error ){_ceg :=font .FontDescriptor ();if _ceg ==nil {return nil ,_gb .New ("\u0063\u006fu\u006c\u0064\u0020\u006e\u006f\u0074\u0020\u0067\u0065\u0074\u0020\u0066\u006f\u006e\u0074\u0020\u0064\u0065\u0073\u0063\u0072\u0069pt\u006f\u0072");};_ccf ,_eced :=_cg .GetStream (_ceg .FontFile2 );if !_eced {return nil ,_gb .New ("\u006di\u0073\u0073\u0069\u006e\u0067\u0020\u0066\u006f\u006e\u0074\u0020f\u0069\u006c\u0065\u0020\u0073\u0074\u0072\u0065\u0061\u006d");};_ee ,_cad :=_cg .DecodeStream (_ccf );if _cad !=nil {return nil ,_cad ;};_ed ,_cad :=_dgc .Parse (_ee );if _cad !=nil {return nil ,_cad ;};if size <=1{size =10;};return &TextFont {Font :font ,Face :_dgc .NewFace (_ed ,&_dgc .Options {Size :size }),Size :size ,_gab :_ed },nil ;};func (_abg *TextFont )WithSize (size float64 ,originalFont *_c .PdfFont )*TextFont {if size <=1{size =10;};return &TextFont {Font :_abg .Font ,Face :_dgc .NewFace (_abg ._gab ,&_dgc .Options {Size :size }),Size :size ,_gab :_abg ._gab ,_ad :originalFont };};const (LineCapRound LineCap =iota ;LineCapButt ;LineCapSquare ;);func (_bdd *TextState )ProcTm (a ,b ,c ,d ,e ,f float64 ){_bdd .Tm =_de .NewMatrix (a ,b ,c ,d ,e ,-f );_bdd .Tlm =_bdd .Tm .Clone ();};func (_dfa *TextState )ProcTD (tx ,ty float64 ){_dfa .Tl =-ty ;_dfa .ProcTd (tx ,ty )};func NewTextState ()*TextState {return &TextState {Th :100,Tm :_de .IdentityMatrix (),Tlm :_de .IdentityMatrix ()};};type Pattern interface{ColorAt (_cc ,_dd int )_d .Color ;};func (_dbg *TextFont )BytesToCharcodes (data []byte )[]_e .CharCode {if _dbg ._ad !=nil {return _dbg ._ad .BytesToCharcodes (data );};return _dbg .Font .BytesToCharcodes (data );};const (FillRuleWinding FillRule =iota ;FillRuleEvenOdd ;);func NewTextFontFromPath (filePath string ,size float64 )(*TextFont ,error ){_af ,_dbb :=_c .NewPdfFontFromTTFFile (filePath );if _dbb !=nil {return nil ,_dbb ;};return NewTextFont (_af ,size );};func (_gcb *TextState )ProcTd (tx ,ty float64 ){_gcb .Tlm .Concat (_de .TranslationMatrix (tx ,-ty ));_gcb .Tm =_gcb .Tlm .Clone ();};func (_gfbb *TextState )Reset (){_gfbb .Tm =_de .IdentityMatrix ();_gfbb .Tlm =_de .IdentityMatrix ()};func (_bbbb *TextState )ProcQ (data []byte ,ctx Context ){_bbbb .ProcTStar ();_bbbb .ProcTj (data ,ctx )};func (_egg *TextState )ProcTj (data []byte ,ctx Context ){_cf :=_egg .Tf .Size ;_fbf :=_egg .Th /100.0;_ffa :=_de .NewMatrix (_cf *_fbf ,0,0,_cf ,0,_egg .Ts );_gcc :=_egg .Tf .CharcodesToUnicode (_egg .Tf .BytesToCharcodes (data ));for _ ,_gce :=range _gcc {if _gce =='\x00'{continue ;};_dfg :=_egg .Tm .Clone ();_egg .Tm .Concat (_ffa );_cfa ,_aac :=_egg .Tm .Transform (0,0);ctx .Scale (1,-1);ctx .DrawString (string (_gce ),_cfa ,_aac );ctx .Scale (1,-1);_aec :=0.0;if _gce ==' '{_aec =_egg .Tw ;};var _aedd float64 ;if _gfb ,_ ,_acg :=_egg .Tf .GetRuneMetrics (_gce );_acg {_aedd =_gfb *0.001*_cf ;}else {_aedd ,_ =ctx .MeasureString (string (_gce ));};_ggg :=(_aedd +_egg .Tc +_aec )*_fbf ;_egg .Tm =_de .TranslationMatrix (_ggg ,0).Mult (_dfg );};};type TextState struct{Tc float64 ;Tw float64 ;Th float64 ;Tl float64 ;Tf *TextFont ;Ts float64 ;Tm _de .Matrix ;Tlm _de .Matrix ;};type LineJoin int ;type FillRule int ;type LineCap int ;func (_gfe *TextState )ProcDQ (data []byte ,aw ,ac float64 ,ctx Context ){_gfe .Tw =aw ;_gfe .Tc =ac ;_gfe .ProcQ (data ,ctx );};func (_ccg *TextState )ProcTf (font *TextFont ){_ccg .Tf =font };func (_bae *TextFont )CharcodesToUnicode (charcodes []_e .CharCode )[]rune {if _bae ._ad !=nil {return _bae ._ad .CharcodesToUnicode (charcodes );};return _bae .Font .CharcodesToUnicode (charcodes );};func (_aed *TextFont )GetCharMetrics (code _e .CharCode )(float64 ,float64 ,bool ){if _gc ,_da :=_aed .Font .GetCharMetrics (code );_da &&_gc .Wx !=0{return _gc .Wx ,_gc .Wy ,_da ;};if _aed ._ad ==nil {return 0,0,false ;};_fa ,_ac :=_aed ._ad .GetCharMetrics (code );return _fa .Wx ,_fa .Wy ,_ac &&_fa .Wx !=0;};type TextFont struct{Font *_c .PdfFont ;Face _ce .Face ;Size float64 ;_gab *_dgc .Font ;_ad *_c .PdfFont ;};func (_dge *TextState )ProcTStar (){_dge .ProcTd (0,-_dge .Tl )};func (_egf *TextFont )GetRuneMetrics (r rune )(float64 ,float64 ,bool ){if _bg ,_dcf :=_egf .Font .GetRuneMetrics (r );_dcf &&_bg .Wx !=0{return _bg .Wx ,_bg .Wy ,_dcf ;};if _egf ._ad ==nil {return 0,0,false ;};_gf ,_ead :=_egf ._ad .GetRuneMetrics (r );return _gf .Wx ,_gf .Wy ,_ead &&_gf .Wx !=0;};const (LineJoinRound LineJoin =iota ;LineJoinBevel ;);type Gradient interface{Pattern ;AddColorStop (_eb float64 ,_f _d .Color );};type Context interface{Push ();Pop ();Matrix ()_de .Matrix ;SetMatrix (_ge _de .Matrix );Translate (_eg ,_dgf float64 );Scale (_b ,_dgg float64 );Rotate (_ca float64 );MoveTo (_def ,_bb float64 );LineTo (_dc ,_geb float64 );CubicTo (_cd ,_fe ,_gg ,_cea ,_ga ,_ea float64 );QuadraticTo (_fd ,_gge ,_ded ,_bc float64 );NewSubPath ();ClosePath ();ClearPath ();Clip ();ClipPreserve ();ResetClip ();LineWidth ()float64 ;SetLineWidth (_ff float64 );SetLineCap (_ec LineCap );SetLineJoin (_ddd LineJoin );SetDash (_gaf ...float64 );SetDashOffset (_cdd float64 );Fill ();FillPreserve ();Stroke ();StrokePreserve ();SetRGBA (_be ,_a ,_bba ,_geg float64 );SetFillRGBA (_gbd ,_bbb ,_ag ,_gef float64 );SetFillStyle (_fb Pattern );SetFillRule (_age FillRule );SetStrokeRGBA (_cda ,_bd ,_dda ,_ece float64 );SetStrokeStyle (_cgd Pattern );TextState ()*TextState ;DrawString (_ddc string ,_aa ,_cca float64 );MeasureString (_ba string )(_dggb ,_dce float64 );DrawRectangle (_db ,_ab ,_eaa ,_bcb float64 );DrawImage (_cce _dg .Image ,_fg ,_cec int );DrawImageAnchored (_defb _dg .Image ,_df ,_eab int ,_ae ,_bda float64 );Height ()int ;Width ()int ;};func (_dfe *TextState )Translate (tx ,ty float64 ){_dfe .Tm =_de .TranslationMatrix (tx ,ty ).Mult (_dfe .Tm );};