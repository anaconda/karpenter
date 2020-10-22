/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package queue

import (
	"github.com/ellistarn/karpenter/pkg/apis/autoscaling/v1alpha1"
	"github.com/ellistarn/karpenter/pkg/cloudprovider"
)

// Producer implements a Pending Capacity metric
type Producer struct {
	*v1alpha1.MetricsProducer
	Queue cloudprovider.Queue
}

// Reconcile of the metrics
func (p *Producer) Reconcile() error {
	_, err := p.Queue.Length()
	if err != nil {
		return err
	}
	return nil
}
