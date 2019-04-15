package appl

import (
	"fmt"
	"testing"
)

func TestFile(t *testing.T) {
	s := `
	
<Publication Version="5.3.1" xmlns="http://ap.org/schemas/03/2005/appl" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
<Identification>
  <ItemId>8d1a9a5830384c6cb76a06a390f24bea</ItemId>
  <RecordId>4e36f5f1f4524e0eb7d0f2818da60b4a</RecordId>
  <CompositeId>27959c2c54c8493ab641d03c987faccb</CompositeId>
  <CompositionType>StandardOnlineVideo</CompositionType>
  <MediaType>Video</MediaType>
  <DefaultLanguage>en</DefaultLanguage>
  <RecordSequenceNumber>0</RecordSequenceNumber>
  <FriendlyKey>4531952946</FriendlyKey>
</Identification>
<PublicationManagement>
  <RecordType>Change</RecordType>
  <FilingType>OnlineVideo</FilingType>
  <IsDistributionReady>true</IsDistributionReady>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <FirstCreated Year="2018" Month="7" Day="3" Time="23:00:25"></FirstCreated>
  <LastModifiedDateTime>2018-07-03T23:00:22</LastModifiedDateTime>
  <Status>Usable</Status>
  <ItemStartDateTime>2018-07-03T23:00:25</ItemStartDateTime>
  <ItemStartDateTimeActual>2018-07-03T23:00:25</ItemStartDateTimeActual>
  <Function>BroadcastVideo</Function>
  <IsDigitized>true</IsDigitized>
</PublicationManagement>
<NewsLines>
  <Title>NewsMinute 5 PM (NR)</Title>
  <HeadLine>AP Top Stories July 3 P</HeadLine>
  <OriginalHeadLine>AP Top Stories July 3 P</OriginalHeadLine>
  <DateLine>Various</DateLine>
  <RightsLine>International: See Script | US: See Script</RightsLine>
  <LocationLine>Various</LocationLine>
</NewsLines>
<AdministrativeMetadata>
  <Provider Type="AP">AP</Provider>
  <Source Type="ThirdParty">Various</Source>
  <TransmissionSource>EapVideo</TransmissionSource>
  <TransmissionSource>NotMemberFeed</TransmissionSource>
  <TransmissionSource>ENPS</TransmissionSource>
  <TransmissionSource>ENPSWAS02</TransmissionSource>
  <TransmissionSource>JagRoars</TransmissionSource>
  <ProductSource>news-us</ProductSource>
  <ContentElement>Package</ContentElement>
  <Reach Scheme="AP">National</Reach>
  <ConsumerReady>true</ConsumerReady>
  <Property Name="EAI:SUBMISSIONPRIORITY"></Property>
  <Property Name="EAI:RECEIVEDFILENAME"></Property>
  <Property Name="EAI:ProducedAspectRatio"></Property>
  <Property Name="TotalDuration"></Property>
</AdministrativeMetadata>
<RightsMetadata></RightsMetadata>
<DescriptiveMetadata>
  <DateLineLocation>
    <LatitudeDD>0.000000000</LatitudeDD>
    <LongitudeDD>0.000000000</LongitudeDD>
  </DateLineLocation>
  <SubjectClassification System="Editorial" Authority="AP Category Code">
    <Occurrence Id="a" Value="National"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="9" Id="f571fdb0894d10048149ba0a2b2ca13e" Value="Special forces" ParentId="62886ddf8d274d3b80c3206f17153281" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="10" Id="62886ddf8d274d3b80c3206f17153281" Value="Armed forces" ParentId="3b7438807d7010048477ba7fa5283c3e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="3b7438807d7010048477ba7fa5283c3e" Value="Military and defense" ParentId="86aad5207dac100488ecba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="86aad5207dac100488ecba7fa5283c3e" Value="Government and politics" ActualMatch="false" TopParent="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="2" Id="75fa814088d710048e89ded1ce465303" Value="Supreme courts" ParentId="16fbfab0893c10048e06ba0a2b2ca13e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="16fbfab0893c10048e06ba0a2b2ca13e" Value="National courts" ParentId="86ba17607dac1004894eba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="86ba17607dac1004894eba7fa5283c3e" Value="Courts" ParentId="86b9d8e07dac1004894dba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="86b9d8e07dac1004894dba7fa5283c3e" Value="Judiciary" ParentId="86aad5207dac100488ecba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="16fbfab0893c10048e06ba0a2b2ca13e" Value="National courts" ParentId="3942e1139d10439a8243b8633edc550e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="3942e1139d10439a8243b8633edc550e" Value="National governments" ParentId="86aad5207dac100488ecba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="24" Id="7cd996a7a95d4550937b427a52a3c054" Value="Higher education" ParentId="1af99ec3cb954ff4b349b32d60d0376d" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="1af99ec3cb954ff4b349b32d60d0376d" Value="Education" ParentId="75a42fd87df7100483eedf092526b43e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="75a42fd87df7100483eedf092526b43e" Value="Social affairs" ActualMatch="false" TopParent="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="3" Id="32f7475887aa10048a5de9f06a797907" Value="Caves" ParentId="8783d248894710048286ba0a2b2ca13e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="8783d248894710048286ba0a2b2ca13e" Value="Environment and nature" ActualMatch="false" TopParent="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="11" Id="0ee28e0697b04b0f84c53dc320d862e0" Value="Navy" ParentId="62886ddf8d274d3b80c3206f17153281" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="4" Id="34d3dad0872a10048d0dd7c656897a02" Value="Affirmative action" ParentId="1b3af533679f46829513eee9c8df6fc5" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="1b3af533679f46829513eee9c8df6fc5" Value="Social diversity" ParentId="08680bf085af10048c4f9a5aeba5fb06" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="08680bf085af10048c4f9a5aeba5fb06" Value="Social issues" ParentId="75a42fd87df7100483eedf092526b43e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="4" Id="34d3dad0872a10048d0dd7c656897a02" Value="Affirmative action" ParentId="008e219885ab10048110ff2260dd383e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="008e219885ab10048110ff2260dd383e" Value="Human rights and civil liberties" ParentId="08680bf085af10048c4f9a5aeba5fb06" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="4" Id="34d3dad0872a10048d0dd7c656897a02" Value="Affirmative action" ParentId="ec28dcdfc4ca4ac9918d3b61427e65c3" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="ec28dcdfc4ca4ac9918d3b61427e65c3" Value="Race and ethnicity" ParentId="08680bf085af10048c4f9a5aeba5fb06" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="4" Id="34d3dad0872a10048d0dd7c656897a02" Value="Affirmative action" ParentId="86bc3e287dac1004895dba7fa5283c3e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="86bc3e287dac1004895dba7fa5283c3e" Value="Political issues" ParentId="86aad5207dac100488ecba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="12" Id="ebf770208985100485c8bc885dbc3010" Value="Museums" ParentId="340f76e5585d48759d22db3e71dc80ef" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="340f76e5585d48759d22db3e71dc80ef" Value="Leisure travel" ParentId="d3d7a339e8c4488f82f050b5d857f1bf" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="d3d7a339e8c4488f82f050b5d857f1bf" Value="Travel" ParentId="3e37e4b87df7100483d5df092526b43e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="3e37e4b87df7100483d5df092526b43e" Value="Lifestyle" ActualMatch="false" TopParent="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="12" Id="ebf770208985100485c8bc885dbc3010" Value="Museums" ParentId="0c78143a0b384d18b575e4c100003330" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="0c78143a0b384d18b575e4c100003330" Value="Recreation and leisure" ParentId="3e37e4b87df7100483d5df092526b43e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="9" Id="e0982188886e100489abcb8225d5863e" Value="Military leadership" ParentId="3b7438807d7010048477ba7fa5283c3e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="15" Id="33bcce9917ac4dd1b9f5770b94916bea" Value="Judicial appointments and nominations" ParentId="a3dae7a888d710048edfded1ce465303" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Id="a3dae7a888d710048edfded1ce465303" Value="Government appointments and nominations" ParentId="86aad5207dac100488ecba7fa5283c3e" ActualMatch="false"></Occurrence>
  </SubjectClassification>
  <SubjectClassification SystemVersion="1" AuthorityVersion="6871" System="Teragram" Authority="AP Subject">
    <Occurrence Count="15" Id="33bcce9917ac4dd1b9f5770b94916bea" Value="Judicial appointments and nominations" ParentId="86b9d8e07dac1004894dba7fa5283c3e" ActualMatch="true"></Occurrence>
  </SubjectClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6886" System="Teragram" Authority="AP Organization">
    <Occurrence Count="5" Id="0592fa0b46784f919459ed587f02462f" Value="U.S. Navy SEALs" ParentId="d85007168f3e43e4a54bd6d62be68d7a" ActualMatch="true"></Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6886" System="Teragram" Authority="AP Organization">
    <Occurrence Id="d85007168f3e43e4a54bd6d62be68d7a" Value="U.S. Navy" ParentId="862d02c54ea746b5b881e2bc901907ce" ActualMatch="false"></Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6886" System="Teragram" Authority="AP Organization">
    <Occurrence Id="862d02c54ea746b5b881e2bc901907ce" Value="United States military" ParentId="898f52e08921100480efba0a2b2ca13e" ActualMatch="false"></Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6886" System="Teragram" Authority="AP Organization">
    <Occurrence Id="898f52e08921100480efba0a2b2ca13e" Value="United States government" ActualMatch="false" TopParent="true"></Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6907" System="Teragram" Authority="AP Party">
    <Occurrence Count="3" Id="a6331b54c46e4ca8b79fa40fada95db5" Value="Donald Trump">
      <Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e"></Property>
      <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6907" System="Teragram" Authority="AP Party">
    <Occurrence Count="1" Id="b45f693e49de4770a5e49ab58bae40e3" Value="Mark Rutte">
      <Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e"></Property>
      <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6907" System="Teragram" Authority="AP Party">
    <Occurrence Count="2" Id="fa91b46b656d4084b50bf384beee608b" Value="Raymond Kethledge">
      <Property Id="5d2fb70d2c364c15968bf7651362602f" Name="PartyType" Value="GOVERNMENT_FIGURE" ParentId="d188b8b8886b100481accb8225d5863e"></Property>
      <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6907" System="Teragram" Authority="AP Party">
    <Occurrence Count="1" Id="b1d550d087874a0393ebfa85dab5ea0a" Value="Barack Obama">
      <Property Id="c9d7fa107e4e1004847adf092526b43e" Name="PartyType" Value="POLITICIAN" ParentId="d188b8b8886b100481accb8225d5863e"></Property>
      <Property Id="d188b8b8886b100481accb8225d5863e" Name="PartyType" Value="PERSON"></Property>
      <Property Id="788b364882f110048df3df092526b43e" Name="AssociatedState" Value="District of Columbia" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Count="14" Id="788b364882f110048df3df092526b43e" Value="District of Columbia" ParentId="661e48387d5b10048291c076b8e3055c" ActualMatch="true">
      <Property Id="9d26a20b35f0484a891740f8189d4c7b" Name="LocationType" Value="City" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="38.91706" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-77.00025" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Id="661e48387d5b10048291c076b8e3055c" Value="United States" ParentId="661850e07d5b100481f7c076b8e3055c" ActualMatch="false">
      <Property Id="01f56e0e654841eca2e69bf2cbcc0526" Name="LocationType" Value="Nation" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="39.76" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-98.5" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Id="661850e07d5b100481f7c076b8e3055c" Value="North America" ActualMatch="false" TopParent="true">
      <Property Id="976d112cd5c3497ea180aeecab922c6b" Name="LocationType" Value="Continent" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="46.07323" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-100.54688" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Count="5" Id="6e92d9b882c7100488e5df092526b43e" Value="Texas" ParentId="661e48387d5b10048291c076b8e3055c" ActualMatch="true">
      <Property Id="0ae5eb8e00e04295a4fc209c94bfe6ef" Name="LocationType" Value="State" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="31.25044" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-99.25061" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Count="12" Id="66219fb07d5b100482e2c076b8e3055c" Value="United Kingdom" ParentId="6618c9f87d5b10048202c076b8e3055c" ActualMatch="true">
      <Property Id="01f56e0e654841eca2e69bf2cbcc0526" Name="LocationType" Value="Nation" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="54" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-4" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Id="6618c9f87d5b10048202c076b8e3055c" Value="Western Europe" ParentId="661850e07d5b100481f4c076b8e3055c" ActualMatch="false">
      <Property Id="424cdfcd69d64fa6869055f7ebf10be4" Name="LocationType" Value="World region" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Id="661850e07d5b100481f4c076b8e3055c" Value="Europe" ActualMatch="false" TopParent="true">
      <Property Id="976d112cd5c3497ea180aeecab922c6b" Name="LocationType" Value="Continent" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="48.69096" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="9.14062" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <EntityClassification SystemVersion="1" AuthorityVersion="6920" System="Teragram" Authority="AP Geography">
    <Occurrence Count="8" Id="71f6ec4882c7100488f1df092526b43e" Value="Austin" ParentId="6e92d9b882c7100488e5df092526b43e" ActualMatch="true">
      <Property Id="9d26a20b35f0484a891740f8189d4c7b" Name="LocationType" Value="City" Permission="Premium"></Property>
      <Property Name="CentroidLatitude" Value="30.26715" Permission="Premium"></Property>
      <Property Name="CentroidLongitude" Value="-97.74306" Permission="Premium"></Property>
    </Occurrence>
  </EntityClassification>
  <AudienceClassification System="Editorial" Authority="AP Audience">
    <Occurrence Id="661E48387D5B10048291C076B8E3055C" Value="United States">
      <Property Id="3446BF8C410D49E59C0A017D8C49F74A" Name="AudienceType" Value="AUDGEOGRAPHY"></Property>
    </Occurrence>
  </AudienceClassification>
  <AudienceClassification System="Editorial" Authority="AP Audience">
    <Occurrence Id="f43adc08760d10048040e6e7a0f4673e" Value="National" ActualMatch="true">
      <Property Id="317C913CF4AA4C5AB9DB610C250B8810" Name="AudienceType" Value="AUDSCOPE"></Property>
    </Occurrence>
  </AudienceClassification>
  <Comment>AP Video US</Comment>
  <Comment>National</Comment>
</DescriptiveMetadata>
<FilingMetadata>
  <Id>a96b5cb4ed384ac2a3a8601cef932185</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<FilingMetadata>
  <Id>ed047806d5264a1cbc89840bfd43b07e</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<FilingMetadata>
  <Id>e1e467a42c244b1f89ae56adaf34bb01</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<FilingMetadata>
  <Id>bfd458d210fc4207bee47a35203173b2</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<FilingMetadata>
  <Id>7bc3736a63994f7a9cd6a1ec4b2b9586</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<FilingMetadata>
  <Id>7f2d66a5f9894b829acc5682295ff225</Id>
  <ArrivalDateTime>2018-07-03T23:00:25</ArrivalDateTime>
  <TransmissionContent>All</TransmissionContent>
  <SlugLine>NewsMinute 5 PM (NR)</SlugLine>
  <Products>
    <Product>100561</Product>
    <Product>45812</Product>
    <Product>45809</Product>
    <Product>45486</Product>
    <Product>45150</Product>
    <Product>44632</Product>
    <Product>43848</Product>
    <Product>43218</Product>
    <Product>10092</Product>
    <Product>6</Product>
    <Product>1</Product>
  </Products>
  <ForeignKeys System="APMediaAcquisition">
    <Keys Id="105" Field="SettingsID"></Keys>
    <Keys Id="XSAN-DC-ENT" Field="SystemDescription"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/MINUTEWORLD%2020180703I_0703T184029.mov" Field="SourceFileName"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="RenameFileName"></Keys>
    <Keys Id="http://10.6.210.130:8080/Playout/Delivery/2018-Week-27/Proxies/MINUTEWORLD%2020180703I_0703T184029.mov.jpg" Field="ThumbnailPath"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APArchive">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemId"></Keys>
    <Keys Id="7cad92c0fa34463fbe1dfeeb2bfedf7c" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="APOnline">
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemId"></Keys>
    <Keys Id="4e36f5f1f4524e0eb7d0f2818da60b4a" Field="RecordId"></Keys>
  </ForeignKeys>
  <ForeignKeys System="Story">
    <Keys Id="apus100153" Field="ID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="MOS">
    <Keys Id="storyfiler.prod.wdc.vap.mos" Field="mosID"></Keys>
    <Keys Id="ENPSWAS02" Field="ncsID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W;6B37DE32-1090-4D33-98984427DC538569" Field="roID"></Keys>
    <Keys Id="ENPSWAS02;P_APVIDEO\W\R_6B37DE32-1090-4D33-98984427DC538569;BA7BD362-EB98-47DF-83C5C5205437D290" Field="storyID"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="BUS.GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS" Field="binaryMosID"></Keys>
    <Keys Id="GALLERY.ORIGIN-ONE-HD.WDC.AP.MOS;d76173cd;800002ec" Field="objID"></Keys>
  </ForeignKeys>
  <ForeignKeys System="mosPayload">
    <Keys Id="1" Field="Approved"></Keys>
    <Keys Id="0" Field="Break"></Keys>
    <Keys Id="JGAINES" Field="Creator"></Keys>
    <Keys Id="0" Field="MediaTime"></Keys>
    <Keys Id="NNASIR" Field="ModBy"></Keys>
    <Keys Id="20180703230022" Field="ModTime"></Keys>
    <Keys Id="80.5" Field="MOSItemDurations"></Keys>
    <Keys Id="MINUTEWORLD 20180703I" Field="MOSObjSlugs"></Keys>
    <Keys Id="NewsMinute 5 PM (NR)-3" Field="MOSSlugs"></Keys>
    <Keys Id="2018-07-03T23:00:22" Field="MOSStoryStatusTime"></Keys>
    <Keys Id="JGAINES" Field="Owner"></Keys>
    <Keys Id="0" Field="pubApproved"></Keys>
    <Keys Id="See Script" Field="Restrictions"></Keys>
    <Keys Id="0" Field="SourceMediaTime"></Keys>
    <Keys Id="JGAINES" Field="SourceModBy"></Keys>
    <Keys Id="20150109220247" Field="SourceModTime"></Keys>
    <Keys Id="0" Field="SourceTextTime"></Keys>
    <Keys Id="Julian" Field="StoryProducer"></Keys>
    <Keys Id="248" Field="TextTime"></Keys>
    <Keys Id="0" Field="LegalBlock"></Keys>
    <Keys Id="ap" Field="apslatebackgroundimage"></Keys>
    <Keys Id="apus100153" Field="APStoryNumber"></Keys>
    <Keys Id="16:9" Field="AspectRatio"></Keys>
    <Keys Id="English/Natsound" Field="AudioDescription"></Keys>
    <Keys Id="Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)" Field="Caption"></Keys>
    <Keys Id="Various - Various" Field="Dateline"></Keys>
    <Keys Id="See Script" Field="DomesticRestrictions"></Keys>
    <Keys Id="National" Field="EditorialCurationUS"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="Headline"></Keys>
    <Keys Id="See Script" Field="InternationalRestrictions"></Keys>
    <Keys Id="f2bfa309099f33ff4ea1e93f20965c13" Field="ItemIdArchive"></Keys>
    <Keys Id="8d1a9a5830384c6cb76a06a390f24bea" Field="ItemIdOnline"></Keys>
    <Keys Id="0" Field="Kill"></Keys>
    <Keys Id="English/Natsound" Field="LanguageType"></Keys>
    <Keys Id="Various" Field="Location"></Keys>
    <Keys Id="16:9" Field="LogicalAspectRatio"></Keys>
    <Keys Id="minuteworld" Field="originalFileName"></Keys>
    <Keys Id="original" Field="ProducedAspectRatio"></Keys>
    <Keys Id="AP Video US" Field="ProductsUS"></Keys>
    <Keys Id="Various" Field="Date"></Keys>
    <Keys Id="Various" Field="Source"></Keys>
    <Keys Id="NR-Package" Field="StoryFormat"></Keys>
    <Keys Id="apus100153" Field="StoryNumber"></Keys>
    <Keys Id="apus" Field="StoryNumberPrefix"></Keys>
    <Keys Id="AP Top Stories July 3 P" Field="StorySummary"></Keys>
    <Keys Id="news-us" Field="StoryWorkgroup"></Keys>
    <Keys Id="3" Field="ENPSItemType"></Keys>
    <Keys Id="False" Field="MetadataOnlyUpdate"></Keys>
  </ForeignKeys>
  <ForeignKeys System="DigView">
    <Keys Id="MINUTEWORLD 20180703I" Field="OriginalFileName"></Keys>
  </ForeignKeys>
</FilingMetadata>
<PublicationComponent Role="Thumbnail" MediaType="Photo">
  <PhotoContentItem Id="1d899bbdbc381f28e20f6a7067001376" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Thumbnail\2018\20180703\23\1d899bbdbc381f28e20f6a7067001376.jpg" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/thumbnail02.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="7769" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_105x070.jpg">
      <Width>105</Width>
      <Height>70</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Caption" MediaType="Text">
  <TextContentItem Id="a81eaa5ff793416b96e3965a2cbcb0b3" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185">
    <DataContent>
      <nitf>
        <body.content>
          <block>
            <p>Here's the latest for Tuesday July 3rd: Trump to rescind Obama-era guidance on affirmative action; Trump talks to 4 possible court nominees; Thai navy SEALs say soccer team healthy; an exhibit devoted to the pleasures of ice cream opens in UK. (July 3)</p>
          </block>
        </body.content>
      </nitf>
    </DataContent>
    <Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="0">
      <Words>45</Words>
    </Characteristics>
  </TextContentItem>
</PublicationComponent>
<PublicationComponent Role="Script" MediaType="Text">
  <TextContentItem Id="4788d306012646748c33a57a5e4b4251" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185">
    <DataContent>
      <nitf>
        <body.content>
          <block>
            <p></p>
            <p></p>
            <p>RESTRICTION SUMMARY/SHOTLIST:</p>
            <p></p>
            <p>(TRUMP TO RESCIND OBAMA-ERA GUIDANCE ON AFFIRMATIVE ACTION)</p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p>Washington -  10 November 2016</p>
            <p>1. FILE PHOTO: President Barack Obama and President-elect Donald Trump shake hands following their meeting in the Oval Office of the White House in Washington, Thursday, Nov. 10, 2016.</p>
            <p></p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p>Cambridge - 03 August 2017 </p>
            <p>2.FILE PHOTO: In this Aug. 30, 2012, file photo, a tour group walks through the campus of Harvard University in Cambridge, Mass.</p>
            <p></p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p> Austin - 27 September 2012</p>
            <p>3. FILE PHOTO: In this Thursday, Sept. 27, 2012 photo, students walk through the University of Texas at Austin campus in Austin, Texas. </p>
            <p></p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p> Lawrenceville - 12 May 2006 </p>
            <p>4. FILE PHOTO: In this Friday, May 12, 2006, file photo, university students in their caps and gowns are silhouetted as they line up for graduation ceremonies in Lawrenceville, N.J.</p>
            <p></p>
            <p></p>
            <p>(TRUMP INTERVIEWS SUPREME COURT NOMINEES)</p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p>Washington - 2 July 2018</p>
            <p>5. President Donald Trump speaking to reporters during a White House meeting with Dutch Prime Minister Mark Rutte.</p>
            <p></p>
            <p>POOL - AP CLIENTS ONLY</p>
            <p>Joint Base Andrews - 3 July 2018</p>
            <p>6. President Donald Trump boarding AFOne</p>
            <p></p>
            <p>FEDERALIST SOCIETY - MANDATORY COURTESY</p>
            <p>Washington - 18 November 2016</p>
            <p>7. Raymond Kethledge, Judge, Sixth Circuit Court of Appeals</p>
            <p></p>
            <p>FEDERALIST SOCIETY - MANDATORY COURTESY</p>
            <p>Washington - 10 March 2018</p>
            <p>8. Amul Thapar, Judge, Sixth Circuit Court of Appeals</p>
            <p></p>
            <p>FEDERALIST SOCIETY - MANDATORY COURTESY</p>
            <p>Washington - 14 November 2013</p>
            <p>9. Brett Kavanaugh, Judge, DC Circuit Court of Appeals:</p>
            <p></p>
            <p>FEDERALIST SOCIETY - MANDATORY COURTESY</p>
            <p>Washington - 10 March 2018</p>
            <p>10. Amy Coney Barrett, Judge, Seventh Circuit Court of Appeals</p>
            <p></p>
            <p></p>
            <p>(THAI NAVY SEALS SAY SOCCER TEAM HEALTHY)</p>
            <p>THAI NAVY SEAL HANDOUT – AP CLIENTS ONLY</p>
            <p>Mae Sai – 2 July 2018 11. Various of boys inside the cave moments after being located by divers </p>
            <p></p>
            <p>ASSOCIATED PRESS – AP CLIENTS ONLY</p>
            <p>Mae Sai – 3 July 2018</p>
            <p>12. Various of news conference Rear Admiral Arpakorn Yookongkaew, Thai Navy SEAL commander speaking</p>
            <p></p>
            <p>(UK ICE CREAM MUSEUM)</p>
            <p>ASSOCIATED PRESS - AP CLIENTS ONLY</p>
            <p>London, UK - 3 July 2018</p>
            <p>13. Close of woman eating ice cream</p>
            <p>14. Various of visitors tasting ice cream at Conehenge cafe</p>
            <p>15. Tilt down of museum worker pouring ice cream</p>
            <p>16. Woman standing under vanilla vapour cloud </p>
            <p>17. Ice cream paraphernalia</p>
            <p>18. Close of ice cream at Conehenge cafe</p>
            <p>19. Woman eating ice cream inside glow-in-the-dark ice cream experience</p>
            <p>20. Recreation of Agnes B. Marshall's Cookery School, visitors tasting ice cream</p>
            <p></p>
            <p></p>
            <p>STORYLINE: </p>
            <p></p>
            <p>THIS AP NEWSMIUTE,</p>
            <p>====================================================================================</p>
            <p>THE TRUMP ADMINISTRATION IS PLANNING TO RESCIND OBAMA-ERA GUIDANCE THAT ENCOURAGED SCHOOLS TO TAKE A PERSON'S RACE INTO ACCOUNT TO INCREASE DIVERSITY IN ADMISSIONS.</p>
            <p></p>
            <p>THE SHIFT WOULD GIVE SCHOOLS AND UNIVERSITIES THE FEDERAL GOVERNMENT'S BLESSING TO TAKE A RACE-NEUTRAL APPROACH TO THE STUDENTS THEY CONSIDER FOR ADMISSION. </p>
            <p></p>
            <p>======================================================================</p>
            <p>PRESIDENT TRUMP PLANS TO CONTINUE MEETING WITH PROSPECTIVE NOMINEES FOR THE SUPREME COURT THROUGH THE WEEK.</p>
            <p></p>
            <p>TRUMP INTERVIEWED FOUR POTENTIAL NOMINEES ON MONDAY, SPENDING ABOUT 45 MINUTES WITH EACH.</p>
            <p></p>
            <p> ALL ARE FEDERAL APPEALS JUDGES: RAYMOND KETHLEDGE, AMUL THAPAR, BRETT KAVANAUGH AND AMY CONEY BARRETT.</p>
            <p></p>
            <p>=====================================================================</p>
            <p>THAI NAVY SEALS SAY ALL 13 PEOPLE TRAPPED IN A FLOODED CAVE IN NORTHERN THAILAND ARE HEALTHY AND BEING LOOKED AFTER BY MEDICS.</p>
            <p></p>
            <p>THE SEAL COMMANDER SAYS THAT SEVEN MEMBERS OF HIS UNIT — INCLUDING A DOCTOR AND A NURSE — ARE NOW WITH THE 12 BOYS AND THEIR COACH IN THE CAVE WHERE THEY TOOK SHELTER AND HAVE BEEN STUCK FOR 10 DAYS.</p>
            <p>==============================================================</p>
            <p>AS LONDONERS BASK IN A LONG HOT SUMMER, AN EXHIBIT WHOLLY DEVOTED TO PLEASURES DERIVED FROM ICE CREAM IS PROMISING TO COOL THE TEMPERATURES.   </p>
            <p></p>
            <p>THE BRITISH MUSEUM OF FOOD OPENS ITS DOORS TO 'SCOOP: A WONDERFUL ICE CREAM WORLD,' A UNIQUE EXHIBIT CELEBRATING THE PAST, PRESENT AND THE FUTURE OF PERHAPS THE WORLD'S BEST LOVED DESSERT.</p>
            <p>================================================================   </p>
            <p>THIS IS JULIAN STYLES, THE ASSOCIATED PRESS, WITH AP NEWSMINUTE.</p>
            <p></p>
            <p></p>
            <p></p>
            <p></p>
            <p>===========================================================</p>
            <p></p>
            <p>Clients are reminded: </p>
            <p>(i) to check the terms of their licence agreements for use of content outside news programming and that further advice and assistance can be obtained from the AP Archive on: Tel +44 (0) 20 7482 7482 Email: info@aparchive.com</p>
            <p>(ii) they should check with the applicable collecting society in their Territory regarding the clearance of any sound recording or performance included within the AP Television News service </p>
            <p>(iii) they have editorial responsibility for the use of all and any content included within the AP Television News service and for libel, privacy, compliance and third party rights applicable to their Territory.</p>
          </block>
        </body.content>
      </nitf>
    </DataContent>
    <Characteristics MimeType="text/xml" Format="NITF" FileExtension="xml" SizeInBytes="1">
      <Words>745</Words>
    </Characteristics>
  </TextContentItem>
</PublicationComponent>
<PublicationComponent Role="Preview" MediaType="Photo">
  <PhotoContentItem Id="18b77183bc381f28e20f6a706700395d" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Preview\2018\20180703\23\18b77183bc381f28e20f6a706700395d.jpg" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/preview01.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="94433" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_400x300.jpg">
      <Width>400</Width>
      <Height>300</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Preview" MediaType="Photo">
  <PhotoContentItem Id="18b7f758bc381f28e20f6a7067006761" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Preview\2018\20180703\23\18b7f758bc381f28e20f6a7067006761.jpg" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/preview.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="98594" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_1920x1080.jpg">
      <Width>1920</Width>
      <Height>1080</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Main" MediaType="Video">
  <VideoContentItem Id="48bb100d2ddf44cab7fa9d51b375dbdb" ArrivedInFilingId="a96b5cb4ed384ac2a3a8601cef932185" BinaryPath="None">
    <Characteristics MimeType="text/plain" Format="IIM" FileExtension="txt" SizeInBytes="0" OriginalFileName="dummy.txt">
      <TotalDuration>81000</TotalDuration>
      <ProducedAspectRatio>original</ProducedAspectRatio>
    </Characteristics>
  </VideoContentItem>
</PublicationComponent>
<PublicationComponent Role="Thumbnail" MediaType="Photo">
  <PhotoContentItem Id="1d89d6edbc381f28e20f6a706700e8d8" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Thumbnail\2018\20180703\23\1d89d6edbc381f28e20f6a706700e8d8.jpg" ArrivedInFilingId="ed047806d5264a1cbc89840bfd43b07e" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/thumbnail.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="7769" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a.jpg">
      <Width>105</Width>
      <Height>70</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Thumbnail" MediaType="Photo">
  <PhotoContentItem Id="1d89c5e1bc381f28e20f6a7067008e34" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Thumbnail\2018\20180703\23\1d89c5e1bc381f28e20f6a7067008e34.jpg" ArrivedInFilingId="e1e467a42c244b1f89ae56adaf34bb01" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/thumbnail03.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="25919" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_180x135.jpg">
      <Width>180</Width>
      <Height>135</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Thumbnail" MediaType="Photo">
  <PhotoContentItem Id="1d897d4bbc381f28e20f6a706700e820" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Thumbnail\2018\20180703\23\1d897d4bbc381f28e20f6a706700e820.jpg" ArrivedInFilingId="bfd458d210fc4207bee47a35203173b2" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/thumbnail01.jpg</BinaryLocation>
    <Characteristics MimeType="image/jpeg" Format="JPEG Baseline" FileExtension="jpg" SizeInBytes="8359" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_092x069.jpg">
      <Width>92</Width>
      <Height>69</Height>
    </Characteristics>
  </PhotoContentItem>
</PublicationComponent>
<PublicationComponent Role="Main" MediaType="Video">
  <VideoContentItem Id="c201299bbc3a1f28e20f6d7034003820" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Main\2018\20180703\23\c201299bbc3a1f28e20f6d7034003820.mp4" ArrivedInFilingId="7bc3736a63994f7a9cd6a1ec4b2b9586" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/secured/x090i.mp4</BinaryLocation>
    <Characteristics MimeType="video/mpeg" Format="MPEG" FileExtension="mp4" SizeInBytes="105681605" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_x090i.mp4">
      <AverageBitRate>10240.000000</AverageBitRate>
      <TotalDuration>81000</TotalDuration>
      <Width>1920</Width>
      <Height>1080</Height>
      <VideoCoder>H.264</VideoCoder>
      <FrameRate>60.000000</FrameRate>
      <SampleRate>44000.000000</SampleRate>
      <ProducedAspectRatio>original</ProducedAspectRatio>
    </Characteristics>
  </VideoContentItem>
</PublicationComponent>
<PublicationComponent Role="Main" MediaType="Video">
  <VideoContentItem Id="8844a9b1bc3b1f28e20f6d70340066f7" Href="\\RQXFILEVIR1.RNDEXT.LOCAL\REPOSITORY\Main\2018\20180703\23\8844a9b1bc3b1f28e20f6d70340066f7.mp4" ArrivedInFilingId="7f2d66a5f9894b829acc5682295ff225" BinaryPath="None">
    <BinaryLocation BinaryPath="Akamai" Sequence="0">None</BinaryLocation>
    <BinaryLocation To="9999-12-31T23:59:59" BinaryPath="URL" Sequence="1">http://mrs.appl.qa.s3.amazonaws.com/8d1a9a5830384c6cb76a06a390f24bea/components/secured/x070n.mp4</BinaryLocation>
    <Characteristics MimeType="video/mpeg" Format="MPEG" FileExtension="mp4" SizeInBytes="5783002" OriginalFileName="4e36f5f1f4524e0eb7d0f2818da60b4a_x070n.mp4">
      <AverageBitRate>448.000000</AverageBitRate>
      <TotalDuration>81000</TotalDuration>
      <Width>576</Width>
      <Height>324</Height>
      <VideoCoder>H.264</VideoCoder>
      <FrameRate>30.000000</FrameRate>
      <SampleRate>44000.000000</SampleRate>
      <ProducedAspectRatio>original</ProducedAspectRatio>
    </Characteristics>
  </VideoContentItem>
</PublicationComponent>
</Publication>`
	jo, _ := XMLToJSON(s)

	fmt.Printf("%s\n", jo.ToString())

}
