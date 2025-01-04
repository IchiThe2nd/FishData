package main

import (
	"fmt"
	"testing"
)

func TestProbes(t *testing.T) {

	t.Run("Create a Probe from further nested", func(t *testing.T) {
		want := "Diva"
		input := `This XML file does not appear to have any style information associated with it. The document tree is shown below.<status software="5.12_8H24" hardware="1.0">
<hostname>Diva</hostname></status>`
		//when
		got := NewSystem(input)
		fmt.Println(got.Hostname)
		//then
		if got.Hostname != want {
			t.Errorf(" got %v wanted %v", got, want)
		}
	})
}

/*
NZOMBIES

N - Null
Z - Zero
O - One
M - Many
B - Boundraries
I - INterfaces
E - Exceptions
S - Seperation

*/

// func TestSomething (t t*testing.T){
// 	t.Run ("when there is an event, it has a status element an attribute of software, test passes") {
// 		// Given

// 		// When
// 		// Then
// 	}
// }

/*
This XML file does not appear to have any style information associated with it. The document tree is shown below.
<status software="5.12_8H24" hardware="1.0">
<hostname>Diva</hostname>
<serial>AC5:66625</serial>
<timezone>-8.00</timezone>
<date>01/03/2025 10:55:57</date>
<power>
<failed>11/14/2024 11:22:42</failed>
<restored>11/14/2024 11:23:01</restored>
</power>
<probes>
<probe>
<name>Tmp</name>
<value>79.4 </value>
<type>Temp</type>
</probe>
<probe>
<name>Dis_pH</name>
<value>8.23 </value>
<type>pH</type>
</probe>
<probe>
<name>ORP</name>
<value>429 </value>
<type>ORP</type>
</probe>
<probe>
<name>Salt</name>
<value>32.6 </value>
<type>Cond</type>
</probe>
<probe>
<name>ReturnA</name>
<value>1.0 </value>
</probe>
<probe>
<name>T5lightsA</name>
<value>1.0 </value>
</probe>
<probe>
<name>TurfScrubberA</name>
<value>0.0 </value>
</probe>
<probe>
<name>Chiller_48A</name>
<value>0.0 </value>
</probe>
<probe>
<name>Co2A</name>
<value>0.0 </value>
</probe>
<probe>
<name>Heaters_2_6A</name>
<value>0.0 </value>
</probe>
<probe>
<name>ACfeedA</name>
<value>0.4 </value>
</probe>
<probe>
<name>Skimmer_8A</name>
<value>0.2 </value>
</probe>
<probe>
<name>ReturnW</name>
<value> 84 </value>
</probe>
<probe>
<name>T5lightsW</name>
<value> 114 </value>
</probe>
<probe>
<name>TurfScrubberW</name>
<value> 1 </value>
</probe>
<probe>
<name>Chiller_48W</name>
<value> 1 </value>
</probe>
<probe>
<name>Co2W</name>
<value> 1 </value>
</probe>
<probe>
<name>Heaters_2_6W</name>
<value> 1 </value>
</probe>
<probe>
<name>ACfeedW</name>
<value> 30 </value>
</probe>
<probe>
<name>Skimmer_8W</name>
<value> 35 </value>
</probe>
<probe>
<name>Volt_2</name>
<value>115 </value>
</probe>
<probe>
<name>Cal_pH</name>
<value>7.21 </value>
<type>pH</type>
</probe>
<probe>
<name>Alkx8</name>
<value>10.05</value>
<type>alk</type>
</probe>
<probe>
<name>Cax8</name>
<value> 489 </value>
<type>ca</type>
</probe>
<probe>
<name>Mgx8</name>
<value>1379 </value>
<type>mg</type>
</probe>
<probe>
<name>Tmpx10</name>
<value>24.4 </value>
<type>Temp</type>
</probe>
<probe>
<name>PARx10</name>
<value> 78 </value>
<type>ASM</type>
</probe>
</probes>
<outlets>
<outlet>
<name>LED_COLOR</name>
<outputID>0</outputID>
<state>TBL</state>
<deviceID>base_Var1</deviceID>
</outlet>
<outlet>
<name>LED_POWER</name>
<outputID>1</outputID>
<state>TBL</state>
<deviceID>base_Var2</deviceID>
</outlet>
<outlet>
<name>T5_InnerActi</name>
<outputID>2</outputID>
<state>TBL</state>
<deviceID>base_Var3</deviceID>
</outlet>
<outlet>
<name>T5_Outer</name>
<outputID>3</outputID>
<state>OFF</state>
<deviceID>base_Var4</deviceID>
</outlet>
<outlet>
<name>SndAlm_I6</name>
<outputID>4</outputID>
<state>AOF</state>
<deviceID>base_Alarm</deviceID>
</outlet>
<outlet>
<name>SndWrn_I7</name>
<outputID>5</outputID>
<state>AOF</state>
<deviceID>base_Warn</deviceID>
</outlet>
<outlet>
<name>EmailAlm_I5</name>
<outputID>6</outputID>
<state>AOF</state>
<deviceID>base_email</deviceID>
</outlet>
<outlet>
<name>Email2Alm_I9</name>
<outputID>7</outputID>
<state>AOF</state>
<deviceID>base_email2</deviceID>
</outlet>
<outlet>
<name>Return</name>
<outputID>8</outputID>
<state>AON</state>
<deviceID>2_1</deviceID>
</outlet>
<outlet>
<name>T5lights</name>
<outputID>9</outputID>
<state>AON</state>
<deviceID>2_2</deviceID>
</outlet>
<outlet>
<name>TurfScrubber</name>
<outputID>10</outputID>
<state>AON</state>
<deviceID>2_3</deviceID>
</outlet>
<outlet>
<name>Chiller_48</name>
<outputID>11</outputID>
<state>AOF</state>
<deviceID>2_4</deviceID>
</outlet>
<outlet>
<name>Co2</name>
<outputID>12</outputID>
<state>AON</state>
<deviceID>2_5</deviceID>
</outlet>
<outlet>
<name>Heaters_2_6</name>
<outputID>13</outputID>
<state>AON</state>
<deviceID>2_6</deviceID>
</outlet>
<outlet>
<name>ACfeed</name>
<outputID>14</outputID>
<state>AON</state>
<deviceID>2_7</deviceID>
</outlet>
<outlet>
<name>Skimmer_8</name>
<outputID>15</outputID>
<state>AON</state>
<deviceID>2_8</deviceID>
</outlet>
<outlet>
<name>SOLENOID2_9</name>
<outputID>16</outputID>
<state>AON</state>
<deviceID>2_9</deviceID>
</outlet>
<outlet>
<name>LinkB_2_10</name>
<outputID>17</outputID>
<state>AON</state>
<deviceID>2_10</deviceID>
</outlet>
<outlet>
<name>ALK_3_1</name>
<outputID>18</outputID>
<state>TBL</state>
<deviceID>3_1</deviceID>
</outlet>
<outlet>
<name>Calc_3_2</name>
<outputID>19</outputID>
<state>TBL</state>
<deviceID>3_2</deviceID>
</outlet>
<outlet>
<name>Cor_5_1</name>
<outputID>21</outputID>
<state>OFF</state>
<deviceID>5_1</deviceID>
</outlet>
<outlet>
<name>TopOff</name>
<outputID>22</outputID>
<state>AOF</state>
<deviceID>6_1</deviceID>
</outlet>
<outlet>
<name>Trident_8_3</name>
<outputID>23</outputID>
<state>AOF</state>
<deviceID>8_3</deviceID>
</outlet>
<outlet>
<name>Alk_8_4</name>
<outputID>24</outputID>
<state>AOF</state>
<deviceID>8_4</deviceID>
</outlet>
<outlet>
<name>AutoFeed</name>
<outputID>25</outputID>
<state>AOF</state>
<deviceID>Cntl_A1</deviceID>
</outlet>
<outlet>
<name>ATO_Cycler</name>
<outputID>26</outputID>
<state>AOF</state>
<deviceID>Cntl_A2</deviceID>
</outlet>
<outlet>
<name>High_Alarm</name>
<outputID>27</outputID>
<state>AOF</state>
<deviceID>Cntl_A3</deviceID>
</outlet>
<outlet>
<name>FLOATS</name>
<outputID>28</outputID>
<state>AOF</state>
<deviceID>Cntl_A4</deviceID>
</outlet>
<outlet>
<name>Feeder_9_1</name>
<outputID>29</outputID>
<state>AOF</state>
<deviceID>9_1</deviceID>
</outlet>
<outlet>
<name>KessilT_11_1</name>
<outputID>30</outputID>
<state>TBL</state>
<deviceID>11_1</deviceID>
</outlet>
<outlet>
<name>Alarm_6_2</name>
<outputID>31</outputID>
<state>AOF</state>
<deviceID>6_2</deviceID>
</outlet>
</outlets>
</status>
*/
