package config

func roundRobin() func() string {

	domains := []string{
		"example.com",
		"example.org",
		"example.net",
		"test.com",
		"mytestsite.com",
		"demo.com",
		"dummy.com",
		"testsite.com",
		"placeholder.com",
		"sampledomain.com",
		"testing123.com",
		"tryme.com",
		"mocksite.com",
		"fakesite.com",
		"mywebsite.com",
		"newdomain.com",
		"demoexample.com",
		"sandbox.com",
		"tryouts.com",
		"testingground.com",
		"experimentsite.com",
		"trialdomain.com",
		"stagingarea.com",
		"localtest.me",
		"devtestsite.com",
		"faketest.com",
		"trydomain.com",
		"betaexample.com",
		"testzone.com",
		"experiments.com",
		"unittest.com",
		"examplenetwork.com",
		"testwebpage.com",
		"webtest.com",
		"domainplaceholder.com",
		"practicepage.com",
		"apitesting.com",
		"webappdemo.com",
		"devtools.com",
		"testlink.com",
		"apitestsite.com",
		"localhost.com",
		"trialwebsite.com",
		"exampleserver.com",
		"testwebserver.com",
		"trialrun.com",
		"openwebtest.com",
		"thedomain.com",
		"mockexample.com",
		"trymeout.com",
	}

	index := 0 // Initial index
	n := len(domains)

	return func() string {
		if n == 0 {
			return "" // Return empty string if the array is empty
		}
		// Select the current domain and update the index
		selectedDomain := domains[index]
		index = (index + 1) % n // Update index in a circular manner
		return selectedDomain
	}
}
