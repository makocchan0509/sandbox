for i in {0..100}
do
    curl http://ELB-masem-27216046.ap-northeast-1.elb.amazonaws.com:8080/outputInfo;
    curl http://ELB-masem-27216046.ap-northeast-1.elb.amazonaws.com:8080/outputError;
    sleep 1;
done
