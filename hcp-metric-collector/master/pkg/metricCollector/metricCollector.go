package metricCollector

import (
	"Hybrid_Cloud/hcp-metric-collector/master/pkg/influx"
	"Hybrid_Cloud/hcp-metric-collector/master/pkg/protobuf"
	"Hybrid_Cloud/util/clusterManager"
	"context"
	"net"

	"google.golang.org/grpc"
	"k8s.io/klog"

	// "openmcp/openmcp/klog"
	// "openmcp/openmcp/openmcp-metric-collector/master/pkg/influx"
	// "openmcp/openmcp/hcp-metric-collector/master/pkg/protobuf"
	// "openmcp/openmcp/util/clusterManager"

	"time"
)

type MetricCollector struct {
	ClusterManager clusterManager.ClusterManager
	Influx         influx.Influx
}

func NewMetricCollector(cm *clusterManager.ClusterManager, INFLUX_IP, INFLUX_PORT, INFLUX_USERNAME, INFLUX_PASSWORD string) *MetricCollector {
	klog.V(4).Info("NewMetricCollector Called")
	mc := &MetricCollector{}
	mc.ClusterManager = *cm
	mc.Influx = *influx.NewInflux(INFLUX_IP, INFLUX_PORT, INFLUX_USERNAME, INFLUX_PASSWORD)

	return mc
}

func (mc *MetricCollector) FindClusterName(data *protobuf.Collection) string {
	klog.V(4).Info("FindClusterName Called")
	IpList := []string{}
	for _, Matricsbatch := range data.Metricsbatchs {
		klog.V(2).Info("[Recieved Data] NodeName: ", Matricsbatch.Node.Name, ", IP: "+Matricsbatch.IP)
		IpList = append(IpList, Matricsbatch.IP)
	}
	clusterName := data.ClusterName
	klog.V(2).Info("=> Recieved Metric Data From '", clusterName, "'")
	return clusterName
}

func (mc *MetricCollector) SendMetrics(ctx context.Context, data *protobuf.Collection) (*protobuf.ReturnValue, error) {
	klog.V(4).Info("SendMetrics Called")

	pTime_start := time.Now()

	clusterName := mc.FindClusterName(data)
	mc.Influx.SaveMetrics(clusterName, data)
	var period_int64 int64

	// openmcpPolicyInstance, target_cluster_policy_err := mc.ClusterManager.Crd_client.OpenMCPPolicy("openmcp").Get("metric-collector-period", metav1.GetOptions{})

	// if target_cluster_policy_err != nil {
	// 	klog.V(0).Info(target_cluster_policy_err)

	// } else {
	// 	a := openmcpPolicyInstance.Spec.Template.Spec.Policies
	// 	period := a[0].Value[0]
	// 	klog.V(3).Info("getPeriodPolicy: ", period+" sec")
	// 	period_int64, _ = strconv.ParseInt(period, 10, 64)
	// }

	klog.V(2).Info("gRPC Return Period")

	pTime_end := time.Since(pTime_start)

	pTime := pTime_end.Seconds()

	return &protobuf.ReturnValue{
		Tick:           period_int64,
		ClusterName:    clusterName,
		ProcessingTime: pTime,
	}, nil
}

func (mc *MetricCollector) StartGRPC(GRPC_PORT string) {
	klog.V(4).Info("StartGRPC Called")
	klog.V(2).Info("Grpc Server Start at Port %s\n", GRPC_PORT)
	l, err := net.Listen("tcp", ":"+GRPC_PORT)
	if err != nil {
		klog.V(0).Info("failed to listen: ", err)

	}
	grpcServer := grpc.NewServer()
	protobuf.RegisterSendMetricsServer(grpcServer, mc)
	if err := grpcServer.Serve(l); err != nil {
		klog.V(0).Info("fail to serve: ", err)

	}

}