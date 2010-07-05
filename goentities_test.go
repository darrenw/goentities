package entities

// Tests substantially copied from the htmlentities.rb rubygem

import "testing"

func checkDecode(output string, input string, t *testing.T){
	decodedout := DecodeString(input)
	if decodedout != output {
		t.Fatalf("Failed to decode %s ; got %s instead of %s",input, decodedout, output)
	}
	return
}

func checkEncode(output string, input string, t *testing.T){
	encodedout,ok := EncodeString(input)
	if !ok || encodedout != output{
		t.Fatalf("Failed to encode %s ; got %s instead of %s",input, encodedout, output)
	}
	return
}

func checkEncodeStyle(output string, input string, style string, t *testing.T){
	encodedout,ok := EncodeStringStyle(input,style)
	if !ok || encodedout != output{
		t.Fatalf("Failed to encode %s ; got %s instead of %s",input,encodedout,output)
	}
	return
}

func checkEncodeSpecStyle(output string, input string, spec string, style string, t *testing.T){
	encodedout,ok := EncodeStringSpecStyle(input,spec,style)
	if !ok || encodedout != output{
		t.Fatalf("Failed to encode %s ; got %s instead of %s",input,encodedout,output)
	}
	return
}

func TestDecodeBasic(t *testing.T){
	checkDecode("&","&amp;",t)
	checkDecode("<","&lt;",t)
	checkDecode("\"","&quot;",t)
	checkDecode(">","&gt;",t)
}

func TestEncodeBasic(t *testing.T){
	checkEncode("&amp;","&",t)
	checkEncode("&lt;","<",t)
	checkEncode("&quot;","\"",t)
	checkEncode("&gt;",">",t)
}

func TestEncodeBasicToDecimal(t *testing.T){
	checkEncodeStyle("&#38;", "&", "decimal", t)
	checkEncodeStyle("&#34;", "\"", "decimal", t)
	checkEncodeStyle("&#60;", "<", "decimal", t)
	checkEncodeStyle("&#62;", ">", "decimal", t)
	checkEncodeStyle("&#39;", "'", "decimal", t)
}

func TestEncodeBasicToHexadecimal(t *testing.T){
	checkEncodeStyle("&#x26;", "&", "hexadecimal", t)
	checkEncodeStyle("&#x22;", "\"", "hexadecimal", t)
	checkEncodeStyle("&#x3c;", "<", "hexadecimal", t)
	checkEncodeStyle("&#x3e;", ">", "hexadecimal", t)
	checkEncodeStyle("&#x27;", "'", "hexadecimal", t)
}

	
func TestEncodeExtendedNamed(t *testing.T){
	checkEncodeStyle("&plusmn;", "±", "named", t)
	checkEncodeStyle("&eth;", "ð", "named", t)
	checkEncodeStyle("&OElig;", "Œ", "named", t)
	checkEncodeStyle("&oelig;", "œ", "named", t)
}

func TestDecodeExtendedNamed(t *testing.T){
	checkDecode("±", "&plusmn;",t)
	checkDecode("ð", "&eth;",t)
	checkDecode("Œ", "&OElig;",t)
	checkDecode("œ", "&oelig;",t);
}
	
func TestDecodeDecimal(t *testing.T){
	checkDecode("“", "&#8220;", t)
	checkDecode("…", "&#8230;", t)
	checkDecode(" ", "&#32;", t)
}

func TestEncodeDecimal(t *testing.T){
	checkEncodeStyle("&#8220;", "“", "decimal", t)
	checkEncodeStyle("&#8230;", "…", "decimal", t)
}

func TestDecodeHexadecimal(t *testing.T){
	checkDecode("−", "&#x2212;", t)
	checkDecode("—", "&#x2014;", t)
	checkDecode("`", "&#x0060;", t)
	checkDecode("`", "&#x60;", t)
}

func TestEncodeHexadecimal(t *testing.T){
	checkEncodeStyle("&#x2212;", "−", "hexadecimal", t)
	checkEncodeStyle("&#x2014;", "—", "hexadecimal", t)
}

func TestDecodeMixedText(t *testing.T){
 checkDecode("Le tabac pourrait bientôt être banni dans tous les lieux publics en France","Le tabac pourrait bient&ocirc;t &#234;tre banni dans tous les lieux publics en France",t)
    checkDecode("\"bientôt\" & 文字","&quot;bient&ocirc;t&quot; &amp; &#25991;&#x5b57;",t)
}

func TestEncodeMixedText(t *testing.T){
	checkEncodeSpecStyle("&quot;bient&ocirc;t&quot; &amp; &#x6587;&#x5b57;","\"bientôt\" & 文字","xhtml1","named-hexadecimal",t) 
	checkEncodeSpecStyle("&quot;bient&ocirc;t&quot; &amp; &#25991;&#23383;","\"bientôt\" & 文字","xhtml1","named",t)
}

func TestDecodeEmptyString(t *testing.T){
	checkDecode("","",t)
}

func TestSkipUnknown(t *testing.T){
	checkDecode("&bogus;","&bogus;",t)
}

func TestDecodeDoubleEncoded(t *testing.T){
	checkDecode("&amp;","&amp;amp;",t)
}

func TestDontEncodeNormalASCII(t *testing.T){
	checkEncodeSpecStyle(" "," ","basic","decimal",t)
	checkEncodeSpecStyle("`","`","basic","decimal",t)
}
