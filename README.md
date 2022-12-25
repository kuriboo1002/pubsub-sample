# pubsub-sample
pub/subのstreamingPull型の検証

## pub/subの作成
export topicName=<topicName>
export projectId=<projectId>

gcloud config set project $projectId
gcloud pubsub topics create $topicName
gcloud pubsub subscriptions create --topic $topicName $topicName-sub

gcloud pubsub topics list
gcloud pubsub subscriptions list

## publish
gcloud pubsub topics publish $topicName --message "PubSub Sample"

## subscription
go run main.go