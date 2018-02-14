# pickpic
Get images from Twitter account.

## Building the command.
```
# Install the goquery.
go get github.com/PuerkitoBio/goquery

# git clone Pickpic
git clone git@github.com:aoimofu/pickpic.git

# Building the command.
cd pickpic/
go build
```

## Usage
```
pickpic -i <a Twitter ID>
```

## example
```
pickpic -i anju_inami

# Output
https://pbs.twimg.com/media/DVwpKmKU0AARmEE.jpg
https://pbs.twimg.com/media/DVjx6XcVAAAQho9.jpg
https://pbs.twimg.com/media/DVjx6XbV4AEvKo4.jpg
https://pbs.twimg.com/media/DVbmt_FVAAEJBcY.jpg
https://pbs.twimg.com/media/DVbFYs6U0AAbLnn.jpg
...( The following is omitted. )
```
