package http_client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetADistributionCenterWithSuccess(t *testing.T) {
	t.Run("Get1DistributionCenter", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"itemId": "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
					"distributionCenters": [
						{
							"id": "66af3398-fd04-471f-aefc-a2280e7d02d3",
							"name": "CD1",
							"zone": "S1",
							"state": "SP",
							"status": "ACTIVE",
							"quantity": 51
						}
					]
				}
				`))
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, dc)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(*dc))
	})

	t.Run("Get1DistributionCenterWithConstructor", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"itemId": "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
					"distributionCenters": [
						{
							"id": "66af3398-fd04-471f-aefc-a2280e7d02d3",
							"name": "CD1",
							"zone": "S1",
							"state": "SP",
							"status": "ACTIVE",
							"quantity": 51
						}
					]
				}
				`))
			}))
		defer server.Close()
		dcRepo := NewDistributionCenterAdapterFactory(server.URL)
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, dc)
		assert.Nil(t, err)
		assert.Equal(t, 1, len(*dc))
	})
	t.Run("Get2DistributionCenter", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
						"itemId": "c60ce040-e2e4-4828-b959-a500996816b8",
						"distributionCenters": [
							{
								"id": "c7f2d3e4-6a9f-1d8b-0c2a-5e4d3b2a7f6c",
								"name": "CD19",
								"zone": "E1",
								"state": "SP",
								"status": "ACTIVE",
								"quantity": 445
							},
							{
								"id": "c7f2d3e4-6a9f-1d8b-0c2a-4d3e5b2a7f6c",
								"name": "CD25",
								"zone": "C1",
								"state": "SP",
								"status": "ACTIVE",
								"quantity": 223
							}
						]
					}
				`))
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("c60ce040-e2e4-4828-b959-a500996816b8")
		assert.NotNil(t, dc)
		assert.Nil(t, err)
		assert.Equal(t, 2, len(*dc))
	})
}

func TestGetADistributionCenterWithErrors(t *testing.T) {
	t.Run("Get1DistributionCenterButHaveAnErrorTryingToDecode", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
					"itemId": "0f99276b-aa53-44f6-8bb6-5b4ededc9615",
					"distributionCenters": [
						{
							"id": "66af3398-fd04-471f-aefc-a2280e7d02d3",
							"name": "CD1",
							"zone": "INVALID_ZONE",
							"state": "SP",
							"status": "ACTIVE",
							"quantity": 51
						}
					]
				}
				`))
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, err)
		assert.Nil(t, dc)
		assert.Equal(t, ErrIntegrationServer, err)
	})
	t.Run("Get1DistributionCenterButGetZeroOfThem", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{
						"itemId": "c60ce040-e2e4-4828-b959-a500996816b8",
						"distributionCenters": []
					}
				`))
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, err)
		assert.Nil(t, dc)
		assert.Equal(t, ErrIntegrationDCsEmpty, err)
	})
	t.Run("Get1DistributionCenterButReceive200WithoutBody", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusOK)
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, err)
		assert.Nil(t, dc)
		assert.Equal(t, ErrIntegrationDecodeJson, err)
	})
	t.Run("Get1DistributionCenterButReceive4XXStatus", func(t *testing.T) {
		server := httptest.NewServer(
			http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, "/distributioncenters", r.URL.Path)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{ 
									"message": "unexpected error"
								}
				`))
			}))
		defer server.Close()
		dcRepo := &DistributionCenterAdapter{server.URL}
		dc, err := dcRepo.GetDCsByItemId("0f99276b-aa53-44f6-8bb6-5b4ededc9615")
		assert.NotNil(t, err)
		assert.Nil(t, dc)
		assert.Equal(t, ErrIntegrationClient, err)
	})
}
