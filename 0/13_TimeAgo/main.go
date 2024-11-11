package main

/*func TimeAgo(pastTime time.Time) string {
	duration := time.Since(pastTime)

	if duration.Seconds() < 60 {
		return fmt.Sprintf("%.0f secundes ago", duration.Seconds())
	} else if duration.Minutes() < 60 {
		return fmt.Sprintf("%.0f minutes ago", duration.Minutes())
	} else if duration.Hours() < 24 {
		return fmt.Sprintf("%.0f hours ago", duration.Hours())
	} else if duration.Hours() < 24*7 {
		return fmt.Sprintf("%.0f дней назад", duration.Hours()/24)
	} else if duration.Hours() < 24*30 {
		return fmt.Sprintf("%.0f недель назад", duration.Hours()/(24*7))
	} else if duration.Hours() < 24*365 {
		return fmt.Sprintf("%.0f month ago", duration.Hours()/(24*30))
	} else {
		return fmt.Sprintf("%.0f yar ago", duration.Hours()/(24*365))
	}
}

func TestTimeAgo(t *testing.T) {
	pastTime := time.Now().Add(-2 * time.Hour)
	timeAgo := TimeAgo(pastTime)
	expected := "2 hours ago"
	if timeAgo != expected {
		t.Errorf("Expected time ago string to be %s, but got %s", expected, timeAgo)
	}
}

func main() {
	pastTime := time.Date(2023, 10, 23, 2, 41, 49, 0, time.UTC)
	result := TimeAgo(pastTime)
	fmt.Println(result) // Output: 1 month ago
}*/
