package controller

var DemoVideos = []Video{
	{
		VideoId:       1,
		User:          DemoUser,
		PlayUrl:       "https://example-bucket123.oss-cn-hangzhou.aliyuncs.com/example-bucket123/%01-20220611170201.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		VideoId:       2,
		User:          DemoUser,
		PlayUrl:       "https://example-bucket123.oss-cn-hangzhou.aliyuncs.com/public/2_VID_20220522_192946.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	UserId:        1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
