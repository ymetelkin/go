package xml

import (
	"fmt"
	"testing"
	"time"
)

func TestXml(t *testing.T) {
	s := `	
	<block>
	<p>Sales of Imcon Edge Connectivity Solutions in India to Commence in 4th Quarter</p> <p>
	  ATLANTA, GA / ACCESSWIRE / September 26, 2019 / <a href="http://pr.report/K6gtGIs4">Imcon International Inc.</a>, the developer of the Internet Backpack, an immediate connectivity solution that allows users to communicate from almost any location on the planet, has commenced sales of its proprietary <a href="http://pr.report/2e0njFdg">Internet Backpack and Edge Connectivity Solutions</a> to South Asia and Oceania through its exclusive reseller <a href="http://pr.report/mz47c0wM">Universal Tree Tech</a>. Over 500 units are expected to be shipped to the region within the next twelve months.
	</p> <p>
	  Today's announcement follows the agreement announced earlier this year that <a href="http://pr.report/f-C3rlqZ">Imcon had named Universal Tree Tech its exclusive distribution partner for ASEAN countries including India and Oceania as well as Australia and New Zealand.</a>
	</p> <p>"3.5 billion people, approximately half of the world's population, does not have internet connectivity, including over 750 million people in India alone," said Rob Loud, CEO of Imcon International. "This agreement brings Imcon another step closer to our goal of connecting these populations and leveling the playing field in terms of access to basic, essential human services including educational opportunities, health care, and social services"</p> <p>
	  Imcon's aggressive growth in this region reflects and supports several regional initiatives by governments to address the digital divide in their countries. These initiatives include <a href="http://pr.report/7SyAkIth">Bharat Broadband Network Limited</a> (Bharatnet), the telecom infrastructure provider established by the government of India which recently committed to connecting an additional 250,000 Indian villages through 2020 that do not currently have access to the internet and Australia's <a href="http://pr.report/QfRpTBri">Universal Services Guarantee</a> program that seeks to connect rural and indigenous communities. It is estimated that Bharatnet will add approximately US$68B per year to India's GDP.
	</p> <p>Randeep Dhillon, CEO, of Universal Tree Tech added, "Universal Tree Tech is thrilled to partner with Imcon to help make universal digital connectivity a reality. "We have already seen the positive effect that the Internet Backpack can have on communities in diverse, underserved and remote locations around the world and expect it to make a profound and long-lasting impact on all the communities we are reaching."</p> <p>This announcement is the latest milestone in Imcon's rapidly expanding global sales and distribution footprint, following new agreements with sales agents in Guatemala (for government and military use) and to Florida-based insurance companies where claims adjusters are leveraging the unique functionality of the Internet Backpack in the field to provide expedited claims assistance to homeowners after natural disasters.</p> <p>
	  Other recent deals include a global agreement with <a href="http://pr.report/2MlcIawG">Orbsat Corp</a> (OTCQB:OSAT) to market an array of immediate connectivity solutions utilizing Orbsat's expertise in satellite-based hardware, services and global sales and distribution capability. Imcon has also launched Imcon Latin America Corp with a regional hub in Costa Rica and appointed <a href="http://pr.report/qVfKu5lY">MDS Seguridad SpA</a> as its exclusive reseller for the Republic of Chile in a deal that also includes non-exclusive rights for all other territories in South America.
	</p> <p>About Universal Tree</p> <p>
	  Universal Tree specializes in development and distribution of Innovative Solutions for the Developing world. The challenges of the developing countries are complex and unique, requiring dynamism and astute integration of technology, products and services. We summon coherence of our vast experience across several verticals integrating technology and a myriad of services including Sales, Distribution, Product Development, Manufacturing and Assembly to build solutions that balance the technology ecosystem while improving the life of the citizens as well as the environment of our planet. Universal Tree enables leading innovative start-ups to scale their business models and introducing their products and services to solve unique problems in green field markets. Please visit <a href="http://pr.report/FltfQIG7">www.universaltree.tech</a> for more information.
	</p> <p>About Imcon International</p> <p>Imcon International, Inc., is an immediate connectivity solutions provider with the ability to provide mobile Internet connections on over 90% of the globe. The Internet Backpack is a remote connectivity solution which allows users to be able to communicate and have computing resources from almost every location on the planet. The Internet Backpack also allows users to create internal wireless networks with large coverage areas. Imcon is developing Edge Connectivity Solutions providing users the ability to harness the power of the Internet in the most remote places and extreme of conditions. Please visit http://imconintl.com for more information.</p> <p>CONTACTS:</p> <p>For Imcon International:</p> <p>Rob Loud</p> <p>Imcon International, Inc.</p> <p>470-210-0760</p> <p>Alan Winnikoff</p> <p>Sayles & Winnikoff Communications</p> <p>212-725-5200 x111</p> <p>For Universal Tree Tech:</p> <p>Randeep Dhillon</p> <p>61 421 715 915</p> <p>SOURCE: Imcon International, Inc.</p>View source version on accesswire.com: <a href="https://www.accesswire.com/561159/Imcon-International-Commences-Sales-and-Shipments-of-Proprietary-Internet-Backpack-to-South-Asia-and-Australia-Through-Exclusive-Reseller-Universal-Tree-Tech">https://www.accesswire.com/561159/Imcon-International-Commences-Sales-and-Shipments-of-Proprietary-Internet-Backpack-to-South-Asia-and-Australia-Through-Exclusive-Reseller-Universal-Tree-Tech</a>
  </block>`

	//s = `<Identification><HeadLine>Vikings, Sage Rosenfels agree to 2-year contract</HeadLine></Identification>`

	start := time.Now()
	nd, err := ParseString(s)
	ts := time.Since(start)

	if err != nil {
		t.Error(err.Error())
	}

	fmt.Printf("%s\n", nd.InlineString())
	fmt.Printf("Duration: %v\n", ts)
}
