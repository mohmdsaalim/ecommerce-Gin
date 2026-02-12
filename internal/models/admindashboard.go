package models

type AdminDashboard struct {
    TotalUsers    int64           `json:"total_users"`
    TotalOrders   int64           `json:"total_orders"`
    TotalProducts int64           `json:"total_products"`
    TotalRevenue  float64         `json:"total_revenue"`
    PendingOrders int64           `json:"pending_orders"`
    RecentOrders  []Order         `json:"recent_orders"`
}