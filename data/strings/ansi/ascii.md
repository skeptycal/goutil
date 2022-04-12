
<div class="page">

<div class="page_content">

<h1>ASCII and ANSI Character Table</h1>

<h3>Content</h3>
<blockquote>
	<a href="#overview">Overview</a><br>
	<a href="#asciicontrol">ASCII Control Characters (0-31 and 127)</a><br>
	<a href="#asciiprint">ASCII Characters6)</a><br>
	<a href="#ansilow">ANSI Characters (128-159)</a><br>
	<a href="#ansihigh">ANSI Characters (160-255)</a><br>
</blockquote>

<h2 id="overview">Overview</h2>

<p>ASCII (American Standard Code for Information Interchange) is a 7-bit character set that contains characters from 0 to 127.</p>

<p>The generic term ANSI (American National Standards Institute) is used for 8-bit character sets. These character sets contain the unchanged ASCII character set. In addition, they contain further characters from 128 to 255, which differ in the various ANSI character sets. There are character sets for western special characters and umlauts, and for Arabic, Greek or Cyrillic characters.</p>

<p>The following table shows which characters are available in which (western) character set:</p>

<div class="hscroll">
    <table class="table_border page_tablecentertext">
    <tbody><tr><th rowspan="2">Character code</th><th rowspan="2">Type</th><th colspan="3">Available in character set</th></tr>
    <tr><th>ASCII</th><th>Windows-1252</th><th>ISO-8859-1</th></tr>
    <tr><td>0 - 31</td><td>Control characters</td><td class="green">YES</td><td class="green">YES</td><td class="green">YES</td></tr>
    <tr><td>32 - 126</td><td>Printable characters</td><td class="green">YES</td><td class="green">YES</td><td class="green">YES</td></tr>
    <tr><td>127</td><td>Control characters</td><td class="green">YES</td><td class="green">YES</td><td class="green">YES</td></tr>
    <tr><td>128 - 159</td><td>Printable characters</td><td class="red">NO</td><td class="green">YES</td><td class="red">NO<br><span class="small gray">(only as control characters)</span></td></tr>
    <tr><td>160 - 255</td><td>Printable characters</td><td class="red">NO</td><td class="green">YES</td><td class="green">YES</td></tr>
    </tbody></table>
</div>


<h2 id="asciicontrol">ASCII Control Characters (0-31 and 127)</h2>

<p>These characters are part of ASCII, Windows-1252 and ISO-8859-1.</p>

<p>The characters with the ASCII codes 0 to 31 and 127 are control characters which are not intended for display.</p>

<p>The caret notation (in column "C") is often used in terminals to display control characters. These can usually be entered using the control key (<kbd>Ctrl</kbd>). For example, the notation "^C" corresponds to the key combination <kbd>Ctrl</kbd>+<kbd>C</kbd>.</p>

<p>The escape sequence (in column "E") is used e.g. in programming languages or search functions to be able to enter control characters as text.</p>

<div class="hscroll">
    <table class="table_border table_stripes page_chartable">
    <tbody><tr><th title="Character" style="text-align: center;">C</th><th title="Decimal character code">Dec</th><th title="Hexadecimal character code">Hex</th><th title="Octal character code">Oct</th><th title="Binary character code">Binary</th><th title="Caret notation">C</th><th title="Escape sequence">E</th><th>Name</th></tr>
    <tr><td class="page_character page_controlchar">NUL</td><td>0</td><td>00</td><td>000</td><td>00000000</td><td>^@</td><td>\0</td><td>Null</td></tr>
    <tr><td class="page_character page_controlchar">SOH</td><td>1</td><td>01</td><td>001</td><td>00000001</td><td>^A</td><td></td><td>Start of Heading</td></tr>
    <tr><td class="page_character page_controlchar">STX</td><td>2</td><td>02</td><td>002</td><td>00000010</td><td>^B</td><td></td><td>Start of Text</td></tr>
    <tr><td class="page_character page_controlchar">ETX</td><td>3</td><td>03</td><td>003</td><td>00000011</td><td>^C</td><td></td><td>End of Text</td></tr>
    <tr><td class="page_character page_controlchar">EOT</td><td>4</td><td>04</td><td>004</td><td>00000100</td><td>^D</td><td></td><td>End of Transmission</td></tr>
    <tr><td class="page_character page_controlchar">ENQ</td><td>5</td><td>05</td><td>005</td><td>00000101</td><td>^E</td><td></td><td>Enquiry</td></tr>
    <tr><td class="page_character page_controlchar">ACK</td><td>6</td><td>06</td><td>006</td><td>00000110</td><td>^F</td><td></td><td>Acknowledge</td></tr>
    <tr><td class="page_character page_controlchar">BEL</td><td>7</td><td>07</td><td>007</td><td>00000111</td><td>^G</td><td>\a</td><td>Bell</td></tr>
    <tr><td class="page_character page_controlchar">BS</td><td>8</td><td>08</td><td>010</td><td>00001000</td><td>^H</td><td>\b</td><td>Backspace</td></tr>
    <tr><td class="page_character page_controlchar">HT</td><td>9</td><td>09</td><td>011</td><td>00001001</td><td>^I</td><td>\t</td><td>Horizontal Tab</td></tr>
    <tr><td class="page_character page_controlchar">LF</td><td>10</td><td>0A</td><td>012</td><td>00001010</td><td>^J</td><td>\n</td><td>Line Feed</td></tr>
    <tr><td class="page_character page_controlchar">VT</td><td>11</td><td>0B</td><td>013</td><td>00001011</td><td>^K</td><td>\v</td><td>Vertical Tab</td></tr>
    <tr><td class="page_character page_controlchar">FF</td><td>12</td><td>0C</td><td>014</td><td>00001100</td><td>^L</td><td>\f</td><td>Form Feed</td></tr>
    <tr><td class="page_character page_controlchar">CR</td><td>13</td><td>0D</td><td>015</td><td>00001101</td><td>^M</td><td>\r</td><td>Carriage Return</td></tr>
    <tr><td class="page_character page_controlchar">SO</td><td>14</td><td>0E</td><td>016</td><td>00001110</td><td>^N</td><td></td><td>Shift Out</td></tr>
    <tr><td class="page_character page_controlchar">SI</td><td>15</td><td>0F</td><td>017</td><td>00001111</td><td>^O</td><td></td><td>Shift In</td></tr>
    <tr><td class="page_character page_controlchar">DLE</td><td>16</td><td>10</td><td>020</td><td>00010000</td><td>^P</td><td></td><td>Data Link Escape</td></tr>
    <tr><td class="page_character page_controlchar">DC1</td><td>17</td><td>11</td><td>021</td><td>00010001</td><td>^Q</td><td></td><td>Device Control 1 (XON)</td></tr>
    <tr><td class="page_character page_controlchar">DC2</td><td>18</td><td>12</td><td>022</td><td>00010010</td><td>^R</td><td></td><td>Device Control 2</td></tr>
    <tr><td class="page_character page_controlchar">DC3</td><td>19</td><td>13</td><td>023</td><td>00010011</td><td>^S</td><td></td><td>Device Control 3 (XOFF)</td></tr>
    <tr><td class="page_character page_controlchar">DC4</td><td>20</td><td>14</td><td>024</td><td>00010100</td><td>^T</td><td></td><td>Device Control 4</td></tr>
    <tr><td class="page_character page_controlchar">NAK</td><td>21</td><td>15</td><td>025</td><td>00010101</td><td>^U</td><td></td><td>Negative Acknowledge</td></tr>
    <tr><td class="page_character page_controlchar">SYN</td><td>22</td><td>16</td><td>026</td><td>00010110</td><td>^V</td><td></td><td>Synchronous Idle</td></tr>
    <tr><td class="page_character page_controlchar">ETB</td><td>23</td><td>17</td><td>027</td><td>00010111</td><td>^W</td><td></td><td>End of Transmission Block</td></tr>
    <tr><td class="page_character page_controlchar">CAN</td><td>24</td><td>18</td><td>030</td><td>00011000</td><td>^X</td><td></td><td>Cancel</td></tr>
    <tr><td class="page_character page_controlchar">EM</td><td>25</td><td>19</td><td>031</td><td>00011001</td><td>^Y</td><td></td><td>End of Medium</td></tr>
    <tr><td class="page_character page_controlchar">SUB</td><td>26</td><td>1A</td><td>032</td><td>00011010</td><td>^Z</td><td></td><td>Substitute</td></tr>
    <tr><td class="page_character page_controlchar">ESC</td><td>27</td><td>1B</td><td>033</td><td>00011011</td><td>^[</td><td>\e</td><td>Escape</td></tr>
    <tr><td class="page_character page_controlchar">FS</td><td>28</td><td>1C</td><td>034</td><td>00011100</td><td>^\</td><td></td><td>File Separator</td></tr>
    <tr><td class="page_character page_controlchar">GS</td><td>29</td><td>1D</td><td>035</td><td>00011101</td><td>^]</td><td></td><td>Group Separator</td></tr>
    <tr><td class="page_character page_controlchar">RS</td><td>30</td><td>1E</td><td>036</td><td>00011110</td><td>^^</td><td></td><td>Record Separator</td></tr>
    <tr><td class="page_character page_controlchar">US</td><td>31</td><td>1F</td><td>037</td><td>00011111</td><td>^_</td><td></td><td>Unit Separator</td></tr>
    <tr><td class="page_character page_controlchar">DEL</td><td>127</td><td>7F</td><td>177</td><td>01111111</td><td>^?</td><td></td><td>Delete</td></tr>
    </tbody></table>
</div>


<h2 id="asciiprint">ASCII Characters (32-126)</h2>

<p>These characters are part of ASCII, Windows-1252 and ISO-8859-1.</p>

<p>Characters with ASCII codes 32 to 126 are so-called printable characters intended for display or output on printers.</p>

<div class="hscroll">
    <table class="table_border table_stripes page_chartable">
    <tbody><tr><th title="Character" style="text-align: center;">C</th><th title="Decimal character code">Dec</th><th title="Hexadecimal character code">Hex</th><th title="Octal character code">Oct</th><th title="Binary character code">Binary</th><th title="Named Character Reference">Entity</th><th>Name</th></tr>
    <tr><td class="page_character">&nbsp;</td><td>32</td><td>20</td><td>040</td><td>00100000</td><td></td><td>Space</td></tr>
    <tr><td class="page_character">!</td><td>33</td><td>21</td><td>041</td><td>00100001</td><td>&amp;excl;</td><td>Exclamation mark</td></tr>
    <tr><td class="page_character">"</td><td>34</td><td>22</td><td>042</td><td>00100010</td><td>&amp;quot;</td><td>Quotation mark</td></tr>
    <tr><td class="page_character">#</td><td>35</td><td>23</td><td>043</td><td>00100011</td><td>&amp;num;</td><td>Number sign</td></tr>
    <tr><td class="page_character">$</td><td>36</td><td>24</td><td>044</td><td>00100100</td><td>&amp;dollar;</td><td>Dollar sign</td></tr>
    <tr><td class="page_character">%</td><td>37</td><td>25</td><td>045</td><td>00100101</td><td>&amp;percnt;</td><td>Percent sign</td></tr>
    <tr><td class="page_character">&amp;</td><td>38</td><td>26</td><td>046</td><td>00100110</td><td>&amp;amp;</td><td>Ampersand</td></tr>
    <tr><td class="page_character">'</td><td>39</td><td>27</td><td>047</td><td>00100111</td><td>&amp;apos;</td><td>Apostrophe</td></tr>
    <tr><td class="page_character">(</td><td>40</td><td>28</td><td>050</td><td>00101000</td><td>&amp;lpar;</td><td>Left parenthesis</td></tr>
    <tr><td class="page_character">)</td><td>41</td><td>29</td><td>051</td><td>00101001</td><td>&amp;rpar;</td><td>Right parenthesis</td></tr>
    <tr><td class="page_character">*</td><td>42</td><td>2A</td><td>052</td><td>00101010</td><td>&amp;ast;</td><td>Asterisk</td></tr>
    <tr><td class="page_character">+</td><td>43</td><td>2B</td><td>053</td><td>00101011</td><td>&amp;plus;</td><td>Plus sign</td></tr>
    <tr><td class="page_character">,</td><td>44</td><td>2C</td><td>054</td><td>00101100</td><td>&amp;comma;</td><td>Comma</td></tr>
    <tr><td class="page_character">-</td><td>45</td><td>2D</td><td>055</td><td>00101101</td><td></td><td>Hyphen (minus sign)</td></tr>
    <tr><td class="page_character">.</td><td>46</td><td>2E</td><td>056</td><td>00101110</td><td>&amp;period;</td><td>Period</td></tr>
    <tr><td class="page_character">/</td><td>47</td><td>2F</td><td>057</td><td>00101111</td><td>&amp;sol;</td><td>Slash</td></tr>
    <tr><td class="page_character">0</td><td>48</td><td>30</td><td>060</td><td>00110000</td><td></td><td>Digit 0</td></tr>
    <tr><td class="page_character">1</td><td>49</td><td>31</td><td>061</td><td>00110001</td><td></td><td>Digit 1</td></tr>
    <tr><td class="page_character">2</td><td>50</td><td>32</td><td>062</td><td>00110010</td><td></td><td>Digit 2</td></tr>
    <tr><td class="page_character">3</td><td>51</td><td>33</td><td>063</td><td>00110011</td><td></td><td>Digit 3</td></tr>
    <tr><td class="page_character">4</td><td>52</td><td>34</td><td>064</td><td>00110100</td><td></td><td>Digit 4</td></tr>
    <tr><td class="page_character">5</td><td>53</td><td>35</td><td>065</td><td>00110101</td><td></td><td>Digit 5</td></tr>
    <tr><td class="page_character">6</td><td>54</td><td>36</td><td>066</td><td>00110110</td><td></td><td>Digit 6</td></tr>
    <tr><td class="page_character">7</td><td>55</td><td>37</td><td>067</td><td>00110111</td><td></td><td>Digit 7</td></tr>
    <tr><td class="page_character">8</td><td>56</td><td>38</td><td>070</td><td>00111000</td><td></td><td>Digit 8</td></tr>
    <tr><td class="page_character">9</td><td>57</td><td>39</td><td>071</td><td>00111001</td><td></td><td>Digit 9</td></tr>
    <tr><td class="page_character">:</td><td>58</td><td>3A</td><td>072</td><td>00111010</td><td>&amp;colon;</td><td>Colon</td></tr>
    <tr><td class="page_character">;</td><td>59</td><td>3B</td><td>073</td><td>00111011</td><td>&amp;semi;</td><td>Semicolon</td></tr>
    <tr><td class="page_character">&lt;</td><td>60</td><td>3C</td><td>074</td><td>00111100</td><td>&amp;lt;</td><td>Less-than</td></tr>
    <tr><td class="page_character">=</td><td>61</td><td>3D</td><td>075</td><td>00111101</td><td>&amp;equals;</td><td>Equals sign</td></tr>
    <tr><td class="page_character">&gt;</td><td>62</td><td>3E</td><td>076</td><td>00111110</td><td>&amp;gt;</td><td>Greater-than</td></tr>
    <tr><td class="page_character">?</td><td>63</td><td>3F</td><td>077</td><td>00111111</td><td>&amp;quest;</td><td>Question mark</td></tr>
    <tr><td class="page_character">@</td><td>64</td><td>40</td><td>100</td><td>01000000</td><td>&amp;commat;</td><td>At sign</td></tr>
    <tr><td class="page_character">A</td><td>65</td><td>41</td><td>101</td><td>01000001</td><td></td><td>Uppercase letter A</td></tr>
    <tr><td class="page_character">B</td><td>66</td><td>42</td><td>102</td><td>01000010</td><td></td><td>Uppercase letter B</td></tr>
    <tr><td class="page_character">C</td><td>67</td><td>43</td><td>103</td><td>01000011</td><td></td><td>Uppercase letter C</td></tr>
    <tr><td class="page_character">D</td><td>68</td><td>44</td><td>104</td><td>01000100</td><td></td><td>Uppercase letter D</td></tr>
    <tr><td class="page_character">E</td><td>69</td><td>45</td><td>105</td><td>01000101</td><td></td><td>Uppercase letter E</td></tr>
    <tr><td class="page_character">F</td><td>70</td><td>46</td><td>106</td><td>01000110</td><td></td><td>Uppercase letter F</td></tr>
    <tr><td class="page_character">G</td><td>71</td><td>47</td><td>107</td><td>01000111</td><td></td><td>Uppercase letter G</td></tr>
    <tr><td class="page_character">H</td><td>72</td><td>48</td><td>110</td><td>01001000</td><td></td><td>Uppercase letter H</td></tr>
    <tr><td class="page_character">I</td><td>73</td><td>49</td><td>111</td><td>01001001</td><td></td><td>Uppercase letter I</td></tr>
    <tr><td class="page_character">J</td><td>74</td><td>4A</td><td>112</td><td>01001010</td><td></td><td>Uppercase letter J</td></tr>
    <tr><td class="page_character">K</td><td>75</td><td>4B</td><td>113</td><td>01001011</td><td></td><td>Uppercase letter K</td></tr>
    <tr><td class="page_character">L</td><td>76</td><td>4C</td><td>114</td><td>01001100</td><td></td><td>Uppercase letter L</td></tr>
    <tr><td class="page_character">M</td><td>77</td><td>4D</td><td>115</td><td>01001101</td><td></td><td>Uppercase letter M</td></tr>
    <tr><td class="page_character">N</td><td>78</td><td>4E</td><td>116</td><td>01001110</td><td></td><td>Uppercase letter N</td></tr>
    <tr><td class="page_character">O</td><td>79</td><td>4F</td><td>117</td><td>01001111</td><td></td><td>Uppercase letter O</td></tr>
    <tr><td class="page_character">P</td><td>80</td><td>50</td><td>120</td><td>01010000</td><td></td><td>Uppercase letter P</td></tr>
    <tr><td class="page_character">Q</td><td>81</td><td>51</td><td>121</td><td>01010001</td><td></td><td>Uppercase letter Q</td></tr>
    <tr><td class="page_character">R</td><td>82</td><td>52</td><td>122</td><td>01010010</td><td></td><td>Uppercase letter R</td></tr>
    <tr><td class="page_character">S</td><td>83</td><td>53</td><td>123</td><td>01010011</td><td></td><td>Uppercase letter S</td></tr>
    <tr><td class="page_character">T</td><td>84</td><td>54</td><td>124</td><td>01010100</td><td></td><td>Uppercase letter T</td></tr>
    <tr><td class="page_character">U</td><td>85</td><td>55</td><td>125</td><td>01010101</td><td></td><td>Uppercase letter U</td></tr>
    <tr><td class="page_character">V</td><td>86</td><td>56</td><td>126</td><td>01010110</td><td></td><td>Uppercase letter V</td></tr>
    <tr><td class="page_character">W</td><td>87</td><td>57</td><td>127</td><td>01010111</td><td></td><td>Uppercase letter W</td></tr>
    <tr><td class="page_character">X</td><td>88</td><td>58</td><td>130</td><td>01011000</td><td></td><td>Uppercase letter X</td></tr>
    <tr><td class="page_character">Y</td><td>89</td><td>59</td><td>131</td><td>01011001</td><td></td><td>Uppercase letter Y</td></tr>
    <tr><td class="page_character">Z</td><td>90</td><td>5A</td><td>132</td><td>01011010</td><td></td><td>Uppercase letter Z</td></tr>
    <tr><td class="page_character">[</td><td>91</td><td>5B</td><td>133</td><td>01011011</td><td>&amp;lbrack;</td><td>Left square bracket</td></tr>
    <tr><td class="page_character">\</td><td>92</td><td>5C</td><td>134</td><td>01011100</td><td>&amp;bsol;</td><td>Backslash</td></tr>
    <tr><td class="page_character">]</td><td>93</td><td>5D</td><td>135</td><td>01011101</td><td>&amp;rsqb;</td><td>Right square bracket</td></tr>
    <tr><td class="page_character">^</td><td>94</td><td>5E</td><td>136</td><td>01011110</td><td>&amp;Hat;</td><td>Caret</td></tr>
    <tr><td class="page_character">_</td><td>95</td><td>5F</td><td>137</td><td>01011111</td><td>&amp;lowbar;</td><td>Underscore</td></tr>
    <tr><td class="page_character">`</td><td>96</td><td>60</td><td>140</td><td>01100000</td><td>&amp;grave;</td><td>Grave accent</td></tr>
    <tr><td class="page_character">a</td><td>97</td><td>61</td><td>141</td><td>01100001</td><td></td><td>Lowercase letter a</td></tr>
    <tr><td class="page_character">b</td><td>98</td><td>62</td><td>142</td><td>01100010</td><td></td><td>Lowercase letter b</td></tr>
    <tr><td class="page_character">c</td><td>99</td><td>63</td><td>143</td><td>01100011</td><td></td><td>Lowercase letter c</td></tr>
    <tr><td class="page_character">d</td><td>100</td><td>64</td><td>144</td><td>01100100</td><td></td><td>Lowercase letter d</td></tr>
    <tr><td class="page_character">e</td><td>101</td><td>65</td><td>145</td><td>01100101</td><td></td><td>Lowercase letter e</td></tr>
    <tr><td class="page_character">f</td><td>102</td><td>66</td><td>146</td><td>01100110</td><td></td><td>Lowercase letter f</td></tr>
    <tr><td class="page_character">g</td><td>103</td><td>67</td><td>147</td><td>01100111</td><td></td><td>Lowercase letter g</td></tr>
    <tr><td class="page_character">h</td><td>104</td><td>68</td><td>150</td><td>01101000</td><td></td><td>Lowercase letter h</td></tr>
    <tr><td class="page_character">i</td><td>105</td><td>69</td><td>151</td><td>01101001</td><td></td><td>Lowercase letter i</td></tr>
    <tr><td class="page_character">j</td><td>106</td><td>6A</td><td>152</td><td>01101010</td><td></td><td>Lowercase letter j</td></tr>
    <tr><td class="page_character">k</td><td>107</td><td>6B</td><td>153</td><td>01101011</td><td></td><td>Lowercase letter k</td></tr>
    <tr><td class="page_character">l</td><td>108</td><td>6C</td><td>154</td><td>01101100</td><td></td><td>Lowercase letter l</td></tr>
    <tr><td class="page_character">m</td><td>109</td><td>6D</td><td>155</td><td>01101101</td><td></td><td>Lowercase letter m</td></tr>
    <tr><td class="page_character">n</td><td>110</td><td>6E</td><td>156</td><td>01101110</td><td></td><td>Lowercase letter n</td></tr>
    <tr><td class="page_character">o</td><td>111</td><td>6F</td><td>157</td><td>01101111</td><td></td><td>Lowercase letter o</td></tr>
    <tr><td class="page_character">p</td><td>112</td><td>70</td><td>160</td><td>01110000</td><td></td><td>Lowercase letter p</td></tr>
    <tr><td class="page_character">q</td><td>113</td><td>71</td><td>161</td><td>01110001</td><td></td><td>Lowercase letter q</td></tr>
    <tr><td class="page_character">r</td><td>114</td><td>72</td><td>162</td><td>01110010</td><td></td><td>Lowercase letter r</td></tr>
    <tr><td class="page_character">s</td><td>115</td><td>73</td><td>163</td><td>01110011</td><td></td><td>Lowercase letter s</td></tr>
    <tr><td class="page_character">t</td><td>116</td><td>74</td><td>164</td><td>01110100</td><td></td><td>Lowercase letter t</td></tr>
    <tr><td class="page_character">u</td><td>117</td><td>75</td><td>165</td><td>01110101</td><td></td><td>Lowercase letter u</td></tr>
    <tr><td class="page_character">v</td><td>118</td><td>76</td><td>166</td><td>01110110</td><td></td><td>Lowercase letter v</td></tr>
    <tr><td class="page_character">w</td><td>119</td><td>77</td><td>167</td><td>01110111</td><td></td><td>Lowercase letter w</td></tr>
    <tr><td class="page_character">x</td><td>120</td><td>78</td><td>170</td><td>01111000</td><td></td><td>Lowercase letter x</td></tr>
    <tr><td class="page_character">y</td><td>121</td><td>79</td><td>171</td><td>01111001</td><td></td><td>Lowercase letter y</td></tr>
    <tr><td class="page_character">z</td><td>122</td><td>7A</td><td>172</td><td>01111010</td><td></td><td>Lowercase letter z</td></tr>
    <tr><td class="page_character">{</td><td>123</td><td>7B</td><td>173</td><td>01111011</td><td>&amp;lbrace;</td><td>Left curly brace</td></tr>
    <tr><td class="page_character">|</td><td>124</td><td>7C</td><td>174</td><td>01111100</td><td>&amp;vert;</td><td>Vertical bar</td></tr>
    <tr><td class="page_character">}</td><td>125</td><td>7D</td><td>175</td><td>01111101</td><td>&amp;rcub;</td><td>Right curly brace</td></tr>
    <tr><td class="page_character">~</td><td>126</td><td>7E</td><td>176</td><td>01111110</td><td></td><td>Tilde</td></tr>
    </tbody></table>
</div>


<h2 id="ansilow">ANSI Characters (128-159)</h2>

<p>These characters are part of Windows-1252. In ISO-8859-1 these characters are control characters.</p>

<div class="hscroll">
    <table class="table_border table_stripes page_chartable">
    <tbody><tr><th title="Character" style="text-align: center;">C</th><th title="Decimal character code">Dec</th><th title="Hexadecimal character code">Hex</th><th title="Octal character code">Oct</th><th title="Binary character code">Binary</th><th title="Named Character Reference">Entity</th><th>Name</th></tr>
    <tr><td class="page_character">€</td><td>128</td><td>80</td><td>200</td><td>10000000</td><td>&amp;euro;</td><td>Euro sign</td></tr>
    <tr><td class="page_character">&nbsp;</td><td>129</td><td>81</td><td>201</td><td>10000001</td><td></td><td>(Not used)</td></tr>
    <tr><td class="page_character">‚</td><td>130</td><td>82</td><td>202</td><td>10000010</td><td>&amp;sbquo;</td><td>Single low-9 quotation mark</td></tr>
    <tr><td class="page_character">ƒ</td><td>131</td><td>83</td><td>203</td><td>10000011</td><td>&amp;fnof;</td><td>Latin small letter f with hook</td></tr>
    <tr><td class="page_character">„</td><td>132</td><td>84</td><td>204</td><td>10000100</td><td>&amp;bdquo;</td><td>Double low-9 quotation mark</td></tr>
    <tr><td class="page_character">…</td><td>133</td><td>85</td><td>205</td><td>10000101</td><td>&amp;hellip;</td><td>Horizontal ellipsis</td></tr>
    <tr><td class="page_character">†</td><td>134</td><td>86</td><td>206</td><td>10000110</td><td>&amp;dagger;</td><td>Dagger</td></tr>
    <tr><td class="page_character">‡</td><td>135</td><td>87</td><td>207</td><td>10000111</td><td>&amp;Dagger;</td><td>Double dagger</td></tr>
    <tr><td class="page_character">ˆ</td><td>136</td><td>88</td><td>210</td><td>10001000</td><td>&amp;circ;</td><td>Modifier letter circumflex accent</td></tr>
    <tr><td class="page_character">‰</td><td>137</td><td>89</td><td>211</td><td>10001001</td><td>&amp;permil;</td><td>Per mille sign</td></tr>
    <tr><td class="page_character">Š</td><td>138</td><td>8A</td><td>212</td><td>10001010</td><td>&amp;Scaron;</td><td>Latin capital letter S with caron</td></tr>
    <tr><td class="page_character">‹</td><td>139</td><td>8B</td><td>213</td><td>10001011</td><td>&amp;lsaquo;</td><td>Single left-pointing angle quotation mark</td></tr>
    <tr><td class="page_character">Œ</td><td>140</td><td>8C</td><td>214</td><td>10001100</td><td>&amp;OElig;</td><td>Latin capital ligature OE</td></tr>
    <tr><td class="page_character">&nbsp;</td><td>141</td><td>8D</td><td>215</td><td>10001101</td><td></td><td>(Not used)</td></tr>
    <tr><td class="page_character">Ž</td><td>142</td><td>8E</td><td>216</td><td>10001110</td><td>&amp;Zcaron;</td><td>Latin capital letter Z with caron</td></tr>
    <tr><td class="page_character">&nbsp;</td><td>143</td><td>8F</td><td>217</td><td>10001111</td><td></td><td>(Not used)</td></tr>
    <tr><td class="page_character">&nbsp;</td><td>144</td><td>90</td><td>220</td><td>10010000</td><td></td><td>(Not used)</td></tr>
    <tr><td class="page_character">‘</td><td>145</td><td>91</td><td>221</td><td>10010001</td><td>&amp;lsquo;</td><td>Left single quotation mark</td></tr>
    <tr><td class="page_character">’</td><td>146</td><td>92</td><td>222</td><td>10010010</td><td>&amp;rsquo;</td><td>Right single quotation mark</td></tr>
    <tr><td class="page_character">“</td><td>147</td><td>93</td><td>223</td><td>10010011</td><td>&amp;ldquo;</td><td>Left double quotation mark</td></tr>
    <tr><td class="page_character">”</td><td>148</td><td>94</td><td>224</td><td>10010100</td><td>&amp;rdquo;</td><td>Right double quotation mark</td></tr>
    <tr><td class="page_character">•</td><td>149</td><td>95</td><td>225</td><td>10010101</td><td>&amp;bull;</td><td>Bullet</td></tr>
    <tr><td class="page_character">–</td><td>150</td><td>96</td><td>226</td><td>10010110</td><td>&amp;ndash;</td><td>En dash</td></tr>
    <tr><td class="page_character">—</td><td>151</td><td>97</td><td>227</td><td>10010111</td><td>&amp;mdash;</td><td>Em dash</td></tr>
    <tr><td class="page_character">˜</td><td>152</td><td>98</td><td>230</td><td>10011000</td><td>&amp;tilde;</td><td>Small tilde</td></tr>
    <tr><td class="page_character">™</td><td>153</td><td>99</td><td>231</td><td>10011001</td><td>&amp;trade;</td><td>Trade mark sign</td></tr>
    <tr><td class="page_character">š</td><td>154</td><td>9A</td><td>232</td><td>10011010</td><td>&amp;scaron;</td><td>Latin small letter s with caron</td></tr>
    <tr><td class="page_character">›</td><td>155</td><td>9B</td><td>233</td><td>10011011</td><td>&amp;rsaquo;</td><td>Single right-pointing angle quotation mark</td></tr>
    <tr><td class="page_character">œ</td><td>156</td><td>9C</td><td>234</td><td>10011100</td><td>&amp;oelig;</td><td>Latin small ligature oe</td></tr>
    <tr><td class="page_character">&nbsp;</td><td>157</td><td>9D</td><td>235</td><td>10011101</td><td></td><td>(Not used)</td></tr>
    <tr><td class="page_character">ž</td><td>158</td><td>9E</td><td>236</td><td>10011110</td><td>&amp;zcaron;</td><td>Latin small letter z with caron</td></tr>
    <tr><td class="page_character">Ÿ</td><td>159</td><td>9F</td><td>237</td><td>10011111</td><td>&amp;Yuml;</td><td>Latin capital letter Y with diaeresis</td></tr>
    </tbody></table>
</div>


<h2 id="ansihigh">ANSI Characters (160-255)</h2>

<p>These characters are part of Windows-1252 and ISO-8859-1.</p>

<div class="hscroll">
    <table class="table_border table_stripes page_chartable">
    <tbody><tr><th title="Character" style="text-align: center;">C</th><th title="Decimal character code">Dec</th><th title="Hexadecimal character code">Hex</th><th title="Octal character code">Oct</th><th title="Binary character code">Binary</th><th title="Named Character Reference">Entity</th><th>Name</th></tr>
    <tr><td class="page_character">&nbsp;</td><td>160</td><td>A0</td><td>240</td><td>10100000</td><td>&amp;nbsp;</td><td>No-break space</td></tr>
    <tr><td class="page_character">¡</td><td>161</td><td>A1</td><td>241</td><td>10100001</td><td>&amp;iexcl;</td><td>Inverted exclamation mark</td></tr>
    <tr><td class="page_character">¢</td><td>162</td><td>A2</td><td>242</td><td>10100010</td><td>&amp;cent;</td><td>Cent sign</td></tr>
    <tr><td class="page_character">£</td><td>163</td><td>A3</td><td>243</td><td>10100011</td><td>&amp;pound;</td><td>Pound sign</td></tr>
    <tr><td class="page_character">¤</td><td>164</td><td>A4</td><td>244</td><td>10100100</td><td>&amp;curren;</td><td>Currency sign</td></tr>
    <tr><td class="page_character">¥</td><td>165</td><td>A5</td><td>245</td><td>10100101</td><td>&amp;yen;</td><td>Yen sign</td></tr>
    <tr><td class="page_character">¦</td><td>166</td><td>A6</td><td>246</td><td>10100110</td><td>&amp;brvbar;</td><td>Broken bar</td></tr>
    <tr><td class="page_character">§</td><td>167</td><td>A7</td><td>247</td><td>10100111</td><td>&amp;sect;</td><td>Section sign</td></tr>
    <tr><td class="page_character">¨</td><td>168</td><td>A8</td><td>250</td><td>10101000</td><td>&amp;DoubleDot;</td><td>Diaeresis</td></tr>
    <tr><td class="page_character">©</td><td>169</td><td>A9</td><td>251</td><td>10101001</td><td>&amp;copy;</td><td>Copyright sign</td></tr>
    <tr><td class="page_character">ª</td><td>170</td><td>AA</td><td>252</td><td>10101010</td><td>&amp;ordf;</td><td>Feminine ordinal indicator</td></tr>
    <tr><td class="page_character">«</td><td>171</td><td>AB</td><td>253</td><td>10101011</td><td>&amp;laquo;</td><td>Left-pointing double angle quotation mark</td></tr>
    <tr><td class="page_character">¬</td><td>172</td><td>AC</td><td>254</td><td>10101100</td><td>&amp;not;</td><td>Not sign</td></tr>
    <tr><td class="page_character">­</td><td>173</td><td>AD</td><td>255</td><td>10101101</td><td>&amp;shy;</td><td>Soft hyphen</td></tr>
    <tr><td class="page_character">®</td><td>174</td><td>AE</td><td>256</td><td>10101110</td><td>&amp;reg;</td><td>Registered sign</td></tr>
    <tr><td class="page_character">¯</td><td>175</td><td>AF</td><td>257</td><td>10101111</td><td>&amp;macr;</td><td>Macron</td></tr>
    <tr><td class="page_character">°</td><td>176</td><td>B0</td><td>260</td><td>10110000</td><td>&amp;deg;</td><td>Degree sign</td></tr>
    <tr><td class="page_character">±</td><td>177</td><td>B1</td><td>261</td><td>10110001</td><td>&amp;plusmn;</td><td>Plus-minus sign</td></tr>
    <tr><td class="page_character">²</td><td>178</td><td>B2</td><td>262</td><td>10110010</td><td>&amp;sup2;</td><td>Superscript two</td></tr>
    <tr><td class="page_character">³</td><td>179</td><td>B3</td><td>263</td><td>10110011</td><td>&amp;sup3;</td><td>Superscript three</td></tr>
    <tr><td class="page_character">´</td><td>180</td><td>B4</td><td>264</td><td>10110100</td><td>&amp;DiacriticalAcute;</td><td>Acute accent</td></tr>
    <tr><td class="page_character">µ</td><td>181</td><td>B5</td><td>265</td><td>10110101</td><td>&amp;micro;</td><td>Micro sign</td></tr>
    <tr><td class="page_character">¶</td><td>182</td><td>B6</td><td>266</td><td>10110110</td><td>&amp;para;</td><td>Pilcrow sign</td></tr>
    <tr><td class="page_character">·</td><td>183</td><td>B7</td><td>267</td><td>10110111</td><td>&amp;CenterDot;</td><td>Middle dot</td></tr>
    <tr><td class="page_character">¸</td><td>184</td><td>B8</td><td>270</td><td>10111000</td><td>&amp;Cedilla;</td><td>Cedilla</td></tr>
    <tr><td class="page_character">¹</td><td>185</td><td>B9</td><td>271</td><td>10111001</td><td>&amp;sup1;</td><td>Superscript one</td></tr>
    <tr><td class="page_character">º</td><td>186</td><td>BA</td><td>272</td><td>10111010</td><td>&amp;ordm;</td><td>Masculine ordinal indicator</td></tr>
    <tr><td class="page_character">»</td><td>187</td><td>BB</td><td>273</td><td>10111011</td><td>&amp;raquo;</td><td>Right-pointing double angle quotation mark</td></tr>
    <tr><td class="page_character">¼</td><td>188</td><td>BC</td><td>274</td><td>10111100</td><td>&amp;frac14;</td><td>Vulgar fraction one quarter</td></tr>
    <tr><td class="page_character">½</td><td>189</td><td>BD</td><td>275</td><td>10111101</td><td>&amp;half;</td><td>Vulgar fraction one half</td></tr>
    <tr><td class="page_character">¾</td><td>190</td><td>BE</td><td>276</td><td>10111110</td><td>&amp;frac34;</td><td>Vulgar fraction three quarters</td></tr>
    <tr><td class="page_character">¿</td><td>191</td><td>BF</td><td>277</td><td>10111111</td><td>&amp;iquest;</td><td>Inverted question mark</td></tr>
    <tr><td class="page_character">À</td><td>192</td><td>C0</td><td>300</td><td>11000000</td><td>&amp;Agrave;</td><td>Latin capital letter A with grave</td></tr>
    <tr><td class="page_character">Á</td><td>193</td><td>C1</td><td>301</td><td>11000001</td><td>&amp;Aacute;</td><td>Latin capital letter A with acute</td></tr>
    <tr><td class="page_character">Â</td><td>194</td><td>C2</td><td>302</td><td>11000010</td><td>&amp;Acirc;</td><td>Latin capital letter A with circumflex</td></tr>
    <tr><td class="page_character">Ã</td><td>195</td><td>C3</td><td>303</td><td>11000011</td><td>&amp;Atilde;</td><td>Latin capital letter A with tilde</td></tr>
    <tr><td class="page_character">Ä</td><td>196</td><td>C4</td><td>304</td><td>11000100</td><td>&amp;Auml;</td><td>Latin capital letter A with diaeresis</td></tr>
    <tr><td class="page_character">Å</td><td>197</td><td>C5</td><td>305</td><td>11000101</td><td>&amp;Aring;</td><td>Latin capital letter A with ring above</td></tr>
    <tr><td class="page_character">Æ</td><td>198</td><td>C6</td><td>306</td><td>11000110</td><td>&amp;AElig;</td><td>Latin capital letter AE</td></tr>
    <tr><td class="page_character">Ç</td><td>199</td><td>C7</td><td>307</td><td>11000111</td><td>&amp;Ccedil;</td><td>Latin capital letter C with cedilla</td></tr>
    <tr><td class="page_character">È</td><td>200</td><td>C8</td><td>310</td><td>11001000</td><td>&amp;Egrave;</td><td>Latin capital letter E with grave</td></tr>
    <tr><td class="page_character">É</td><td>201</td><td>C9</td><td>311</td><td>11001001</td><td>&amp;Eacute;</td><td>Latin capital letter E with acute</td></tr>
    <tr><td class="page_character">Ê</td><td>202</td><td>CA</td><td>312</td><td>11001010</td><td>&amp;Ecirc;</td><td>Latin capital letter E with circumflex</td></tr>
    <tr><td class="page_character">Ë</td><td>203</td><td>CB</td><td>313</td><td>11001011</td><td>&amp;Euml;</td><td>Latin capital letter E with diaeresis</td></tr>
    <tr><td class="page_character">Ì</td><td>204</td><td>CC</td><td>314</td><td>11001100</td><td>&amp;Igrave;</td><td>Latin capital letter I with grave</td></tr>
    <tr><td class="page_character">Í</td><td>205</td><td>CD</td><td>315</td><td>11001101</td><td>&amp;Iacute;</td><td>Latin capital letter I with acute</td></tr>
    <tr><td class="page_character">Î</td><td>206</td><td>CE</td><td>316</td><td>11001110</td><td>&amp;Icirc;</td><td>Latin capital letter I with circumflex</td></tr>
    <tr><td class="page_character">Ï</td><td>207</td><td>CF</td><td>317</td><td>11001111</td><td>&amp;Iuml;</td><td>Latin capital letter I with diaeresis</td></tr>
    <tr><td class="page_character">Ð</td><td>208</td><td>D0</td><td>320</td><td>11010000</td><td>&amp;ETH;</td><td>Latin capital letter Eth</td></tr>
    <tr><td class="page_character">Ñ</td><td>209</td><td>D1</td><td>321</td><td>11010001</td><td>&amp;Ntilde;</td><td>Latin capital letter N with tilde</td></tr>
    <tr><td class="page_character">Ò</td><td>210</td><td>D2</td><td>322</td><td>11010010</td><td>&amp;Ograve;</td><td>Latin capital letter O with grave</td></tr>
    <tr><td class="page_character">Ó</td><td>211</td><td>D3</td><td>323</td><td>11010011</td><td>&amp;Oacute;</td><td>Latin capital letter O with acute</td></tr>
    <tr><td class="page_character">Ô</td><td>212</td><td>D4</td><td>324</td><td>11010100</td><td>&amp;Ocirc;</td><td>Latin capital letter O with circumflex</td></tr>
    <tr><td class="page_character">Õ</td><td>213</td><td>D5</td><td>325</td><td>11010101</td><td>&amp;Otilde;</td><td>Latin capital letter O with tilde</td></tr>
    <tr><td class="page_character">Ö</td><td>214</td><td>D6</td><td>326</td><td>11010110</td><td>&amp;Ouml;</td><td>Latin capital letter O with diaeresis</td></tr>
    <tr><td class="page_character">×</td><td>215</td><td>D7</td><td>327</td><td>11010111</td><td>&amp;times;</td><td>Multiplication sign</td></tr>
    <tr><td class="page_character">Ø</td><td>216</td><td>D8</td><td>330</td><td>11011000</td><td>&amp;Oslash;</td><td>Latin capital letter O with stroke</td></tr>
    <tr><td class="page_character">Ù</td><td>217</td><td>D9</td><td>331</td><td>11011001</td><td>&amp;Ugrave;</td><td>Latin capital letter U with grave</td></tr>
    <tr><td class="page_character">Ú</td><td>218</td><td>DA</td><td>332</td><td>11011010</td><td>&amp;Uacute;</td><td>Latin capital letter U with acute</td></tr>
    <tr><td class="page_character">Û</td><td>219</td><td>DB</td><td>333</td><td>11011011</td><td>&amp;Ucirc;</td><td>Latin capital letter U with circumflex</td></tr>
    <tr><td class="page_character">Ü</td><td>220</td><td>DC</td><td>334</td><td>11011100</td><td>&amp;Uuml;</td><td>Latin capital letter U with diaeresis</td></tr>
    <tr><td class="page_character">Ý</td><td>221</td><td>DD</td><td>335</td><td>11011101</td><td>&amp;Yacute;</td><td>Latin capital letter Y with acute</td></tr>
    <tr><td class="page_character">Þ</td><td>222</td><td>DE</td><td>336</td><td>11011110</td><td>&amp;THORN;</td><td>Latin capital letter Thorn</td></tr>
    <tr><td class="page_character">ß</td><td>223</td><td>DF</td><td>337</td><td>11011111</td><td>&amp;szlig;</td><td>Latin small letter sharp s</td></tr>
    <tr><td class="page_character">à</td><td>224</td><td>E0</td><td>340</td><td>11100000</td><td>&amp;agrave;</td><td>Latin small letter a with grave</td></tr>
    <tr><td class="page_character">á</td><td>225</td><td>E1</td><td>341</td><td>11100001</td><td>&amp;aacute;</td><td>Latin small letter a with acute</td></tr>
    <tr><td class="page_character">â</td><td>226</td><td>E2</td><td>342</td><td>11100010</td><td>&amp;acirc;</td><td>Latin small letter a with circumflex</td></tr>
    <tr><td class="page_character">ã</td><td>227</td><td>E3</td><td>343</td><td>11100011</td><td>&amp;atilde;</td><td>Latin small letter a with tilde</td></tr>
    <tr><td class="page_character">ä</td><td>228</td><td>E4</td><td>344</td><td>11100100</td><td>&amp;auml;</td><td>Latin small letter a with diaeresis</td></tr>
    <tr><td class="page_character">å</td><td>229</td><td>E5</td><td>345</td><td>11100101</td><td>&amp;aring;</td><td>Latin small letter a with ring above</td></tr>
    <tr><td class="page_character">æ</td><td>230</td><td>E6</td><td>346</td><td>11100110</td><td>&amp;aelig;</td><td>Latin small letter ae</td></tr>
    <tr><td class="page_character">ç</td><td>231</td><td>E7</td><td>347</td><td>11100111</td><td>&amp;ccedil;</td><td>Latin small letter c with cedilla</td></tr>
    <tr><td class="page_character">è</td><td>232</td><td>E8</td><td>350</td><td>11101000</td><td>&amp;egrave;</td><td>Latin small letter e with grave</td></tr>
    <tr><td class="page_character">é</td><td>233</td><td>E9</td><td>351</td><td>11101001</td><td>&amp;eacute;</td><td>Latin small letter e with acute</td></tr>
    <tr><td class="page_character">ê</td><td>234</td><td>EA</td><td>352</td><td>11101010</td><td>&amp;ecirc;</td><td>Latin small letter e with circumflex</td></tr>
    <tr><td class="page_character">ë</td><td>235</td><td>EB</td><td>353</td><td>11101011</td><td>&amp;euml;</td><td>Latin small letter e with diaeresis</td></tr>
    <tr><td class="page_character">ì</td><td>236</td><td>EC</td><td>354</td><td>11101100</td><td>&amp;igrave;</td><td>Latin small letter i with grave</td></tr>
    <tr><td class="page_character">í</td><td>237</td><td>ED</td><td>355</td><td>11101101</td><td>&amp;iacute;</td><td>Latin small letter i with acute</td></tr>
    <tr><td class="page_character">î</td><td>238</td><td>EE</td><td>356</td><td>11101110</td><td>&amp;icirc;</td><td>Latin small letter i with circumflex</td></tr>
    <tr><td class="page_character">ï</td><td>239</td><td>EF</td><td>357</td><td>11101111</td><td>&amp;iuml;</td><td>Latin small letter i with diaeresis</td></tr>
    <tr><td class="page_character">ð</td><td>240</td><td>F0</td><td>360</td><td>11110000</td><td>&amp;eth;</td><td>Latin small letter eth</td></tr>
    <tr><td class="page_character">ñ</td><td>241</td><td>F1</td><td>361</td><td>11110001</td><td>&amp;ntilde;</td><td>Latin small letter n with tilde</td></tr>
    <tr><td class="page_character">ò</td><td>242</td><td>F2</td><td>362</td><td>11110010</td><td>&amp;ograve;</td><td>Latin small letter o with grave</td></tr>
    <tr><td class="page_character">ó</td><td>243</td><td>F3</td><td>363</td><td>11110011</td><td>&amp;oacute;</td><td>Latin small letter o with acute</td></tr>
    <tr><td class="page_character">ô</td><td>244</td><td>F4</td><td>364</td><td>11110100</td><td>&amp;ocirc;</td><td>Latin small letter o with circumflex</td></tr>
    <tr><td class="page_character">õ</td><td>245</td><td>F5</td><td>365</td><td>11110101</td><td>&amp;otilde;</td><td>Latin small letter o with tilde</td></tr>
    <tr><td class="page_character">ö</td><td>246</td><td>F6</td><td>366</td><td>11110110</td><td>&amp;ouml;</td><td>Latin small letter o with diaeresis</td></tr>
    <tr><td class="page_character">÷</td><td>247</td><td>F7</td><td>367</td><td>11110111</td><td>&amp;divide;</td><td>Division sign</td></tr>
    <tr><td class="page_character">ø</td><td>248</td><td>F8</td><td>370</td><td>11111000</td><td>&amp;oslash;</td><td>Latin small letter o with stroke</td></tr>
    <tr><td class="page_character">ù</td><td>249</td><td>F9</td><td>371</td><td>11111001</td><td>&amp;ugrave;</td><td>Latin small letter u with grave</td></tr>
    <tr><td class="page_character">ú</td><td>250</td><td>FA</td><td>372</td><td>11111010</td><td>&amp;uacute;</td><td>Latin small letter u with acute</td></tr>
    <tr><td class="page_character">û</td><td>251</td><td>FB</td><td>373</td><td>11111011</td><td>&amp;ucirc;</td><td>Latin small letter u with circumflex</td></tr>
    <tr><td class="page_character">ü</td><td>252</td><td>FC</td><td>374</td><td>11111100</td><td>&amp;uuml;</td><td>Latin small letter u with diaeresis</td></tr>
    <tr><td class="page_character">ý</td><td>253</td><td>FD</td><td>375</td><td>11111101</td><td>&amp;yacute;</td><td>Latin small letter y with acute</td></tr>
    <tr><td class="page_character">þ</td><td>254</td><td>FE</td><td>376</td><td>11111110</td><td>&amp;thorn;</td><td>Latin small letter thorn</td></tr>
    <tr><td class="page_character">ÿ</td><td>255</td><td>FF</td><td>377</td><td>11111111</td><td>&amp;yuml;</td><td>Latin small letter y with diaeresis</td></tr>
    </tbody></table>
</div>
