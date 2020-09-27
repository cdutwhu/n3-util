package n3xml

import "testing"

func TestRmComment(t *testing.T) {
	xml := `<WellbeingPersonLink RefId="2FFB63B4-CFEF-4820-8501-E7D1E54555CB">
	<WellbeingEventRefId>D3E34B35-9D75-101A-8C3D-00AA001A1688</WellbeingEventRefId>
	<WellbeingResponseRefId>E6E34B35-9D75-101A-8C3D-00AA001A1700</WellbeingResponseRefId>
	<GroupId>Inv0009</GroupId>
	<PersonRefId SIF_RefObject="StudentPersonal">D3E34B35-9D75-101A-8C3D-00AA001A1652</PersonRefId>


	<!-- Should you have the school here??? <!--
	and also this	-->
	<SchoolInfoRefId>CA285746-359D-7510-1A8C-36432A901A16</SchoolInfoRefId> -->

	
	<HowInvolved>Victim</HowInvolved>
	<FollowUpActionList>
	 <FollowUpAction>
		 <FollowUpDetails>Parent's notified</FollowUpDetails>
	 </FollowUpAction>
	</FollowUpActionList>
</WellbeingPersonLink>`
	xml = RmComment(xml)
	fPln(xml)
}
