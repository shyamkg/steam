// DON'T EDIT *** generated by scaneo *** DON'T EDIT //

package data

import "database/sql"

func ScanMeta(r *sql.Row) (Meta, error) {
	var s Meta
	if err := r.Scan(
		&s.Id,
		&s.Key,
		&s.Value,
	); err != nil {
		return Meta{}, err
	}
	return s, nil
}

func ScanMetas(rs *sql.Rows) ([]Meta, error) {
	structs := make([]Meta, 0, 16)
	var err error
	for rs.Next() {
		var s Meta
		if err = rs.Scan(
			&s.Id,
			&s.Key,
			&s.Value,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEntityHistory(r *sql.Row) (EntityHistory, error) {
	var s EntityHistory
	if err := r.Scan(
		&s.IdentityId,
		&s.Action,
		&s.Description,
		&s.Created,
	); err != nil {
		return EntityHistory{}, err
	}
	return s, nil
}

func ScanEntityHistorys(rs *sql.Rows) ([]EntityHistory, error) {
	structs := make([]EntityHistory, 0, 16)
	var err error
	for rs.Next() {
		var s EntityHistory
		if err = rs.Scan(
			&s.IdentityId,
			&s.Action,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPrivilege(r *sql.Row) (Privilege, error) {
	var s Privilege
	if err := r.Scan(
		&s.Type,
		&s.WorkgroupId,
		&s.EntityType,
		&s.EntityId,
	); err != nil {
		return Privilege{}, err
	}
	return s, nil
}

func ScanPrivileges(rs *sql.Rows) ([]Privilege, error) {
	structs := make([]Privilege, 0, 16)
	var err error
	for rs.Next() {
		var s Privilege
		if err = rs.Scan(
			&s.Type,
			&s.WorkgroupId,
			&s.EntityType,
			&s.EntityId,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEntityPrivilege(r *sql.Row) (EntityPrivilege, error) {
	var s EntityPrivilege
	if err := r.Scan(
		&s.Type,
		&s.WorkgroupId,
		&s.WorkgroupName,
		&s.WorkgroupDescription,
	); err != nil {
		return EntityPrivilege{}, err
	}
	return s, nil
}

func ScanEntityPrivileges(rs *sql.Rows) ([]EntityPrivilege, error) {
	structs := make([]EntityPrivilege, 0, 16)
	var err error
	for rs.Next() {
		var s EntityPrivilege
		if err = rs.Scan(
			&s.Type,
			&s.WorkgroupId,
			&s.WorkgroupName,
			&s.WorkgroupDescription,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanPermission(r *sql.Row) (Permission, error) {
	var s Permission
	if err := r.Scan(
		&s.Id,
		&s.Code,
		&s.Description,
	); err != nil {
		return Permission{}, err
	}
	return s, nil
}

func ScanPermissions(rs *sql.Rows) ([]Permission, error) {
	structs := make([]Permission, 0, 16)
	var err error
	for rs.Next() {
		var s Permission
		if err = rs.Scan(
			&s.Id,
			&s.Code,
			&s.Description,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEntityType(r *sql.Row) (EntityType, error) {
	var s EntityType
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return EntityType{}, err
	}
	return s, nil
}

func ScanEntityTypes(rs *sql.Rows) ([]EntityType, error) {
	structs := make([]EntityType, 0, 16)
	var err error
	for rs.Next() {
		var s EntityType
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRole(r *sql.Row) (Role, error) {
	var s Role
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Role{}, err
	}
	return s, nil
}

func ScanRoles(rs *sql.Rows) ([]Role, error) {
	structs := make([]Role, 0, 16)
	var err error
	for rs.Next() {
		var s Role
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanWorkgroup(r *sql.Row) (Workgroup, error) {
	var s Workgroup
	if err := r.Scan(
		&s.Id,
		&s.Type,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Workgroup{}, err
	}
	return s, nil
}

func ScanWorkgroups(rs *sql.Rows) ([]Workgroup, error) {
	structs := make([]Workgroup, 0, 16)
	var err error
	for rs.Next() {
		var s Workgroup
		if err = rs.Scan(
			&s.Id,
			&s.Type,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentity(r *sql.Row) (Identity, error) {
	var s Identity
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.IsActive,
		&s.LastLogin,
		&s.Created,
	); err != nil {
		return Identity{}, err
	}
	return s, nil
}

func ScanIdentitys(rs *sql.Rows) ([]Identity, error) {
	structs := make([]Identity, 0, 16)
	var err error
	for rs.Next() {
		var s Identity
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.IsActive,
			&s.LastLogin,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentityAndPassword(r *sql.Row) (IdentityAndPassword, error) {
	var s IdentityAndPassword
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Password,
		&s.WorkgroupId,
		&s.IsActive,
		&s.LastLogin,
		&s.Created,
	); err != nil {
		return IdentityAndPassword{}, err
	}
	return s, nil
}

func ScanIdentityAndPasswords(rs *sql.Rows) ([]IdentityAndPassword, error) {
	structs := make([]IdentityAndPassword, 0, 16)
	var err error
	for rs.Next() {
		var s IdentityAndPassword
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Password,
			&s.WorkgroupId,
			&s.IsActive,
			&s.LastLogin,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanIdentityAndRole(r *sql.Row) (IdentityAndRole, error) {
	var s IdentityAndRole
	if err := r.Scan(
		&s.Kind,
		&s.IdentityId,
		&s.IdentityName,
		&s.RoleId,
		&s.RoleName,
	); err != nil {
		return IdentityAndRole{}, err
	}
	return s, nil
}

func ScanIdentityAndRoles(rs *sql.Rows) ([]IdentityAndRole, error) {
	structs := make([]IdentityAndRole, 0, 16)
	var err error
	for rs.Next() {
		var s IdentityAndRole
		if err = rs.Scan(
			&s.Kind,
			&s.IdentityId,
			&s.IdentityName,
			&s.RoleId,
			&s.RoleName,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanEngine(r *sql.Row) (Engine, error) {
	var s Engine
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Location,
		&s.EngineType,
		&s.Created,
	); err != nil {
		return Engine{}, err
	}
	return s, nil
}

func ScanEngines(rs *sql.Rows) ([]Engine, error) {
	structs := make([]Engine, 0, 16)
	var err error
	for rs.Next() {
		var s Engine
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Location,
			&s.EngineType,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanClusterType(r *sql.Row) (ClusterType, error) {
	var s ClusterType
	if err := r.Scan(
		&s.Id,
		&s.Name,
	); err != nil {
		return ClusterType{}, err
	}
	return s, nil
}

func ScanClusterTypes(rs *sql.Rows) ([]ClusterType, error) {
	structs := make([]ClusterType, 0, 16)
	var err error
	for rs.Next() {
		var s ClusterType
		if err = rs.Scan(
			&s.Id,
			&s.Name,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanCluster(r *sql.Row) (Cluster, error) {
	var s Cluster
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.TypeId,
		&s.DetailId,
		&s.Address,
		&s.Messaging,
		&s.State,
		&s.Created,
	); err != nil {
		return Cluster{}, err
	}
	return s, nil
}

func ScanClusters(rs *sql.Rows) ([]Cluster, error) {
	structs := make([]Cluster, 0, 16)
	var err error
	for rs.Next() {
		var s Cluster
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.TypeId,
			&s.DetailId,
			&s.Address,
			&s.Messaging,
			&s.State,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanYarnCluster(r *sql.Row) (YarnCluster, error) {
	var s YarnCluster
	if err := r.Scan(
		&s.Id,
		&s.EngineId,
		&s.EngineType,
		&s.Size,
		&s.ApplicationId,
		&s.Memory,
		&s.Username,
		&s.OutputDir,
	); err != nil {
		return YarnCluster{}, err
	}
	return s, nil
}

func ScanYarnClusters(rs *sql.Rows) ([]YarnCluster, error) {
	structs := make([]YarnCluster, 0, 16)
	var err error
	for rs.Next() {
		var s YarnCluster
		if err = rs.Scan(
			&s.Id,
			&s.EngineId,
			&s.EngineType,
			&s.Size,
			&s.ApplicationId,
			&s.Memory,
			&s.Username,
			&s.OutputDir,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanProject(r *sql.Row) (Project, error) {
	var s Project
	if err := r.Scan(
		&s.Id,
		&s.Name,
		&s.Description,
		&s.ModelCategory,
		&s.Created,
	); err != nil {
		return Project{}, err
	}
	return s, nil
}

func ScanProjects(rs *sql.Rows) ([]Project, error) {
	structs := make([]Project, 0, 16)
	var err error
	for rs.Next() {
		var s Project
		if err = rs.Scan(
			&s.Id,
			&s.Name,
			&s.Description,
			&s.ModelCategory,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanDatasource(r *sql.Row) (Datasource, error) {
	var s Datasource
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.Name,
		&s.Description,
		&s.Kind,
		&s.Configuration,
		&s.Created,
	); err != nil {
		return Datasource{}, err
	}
	return s, nil
}

func ScanDatasources(rs *sql.Rows) ([]Datasource, error) {
	structs := make([]Datasource, 0, 16)
	var err error
	for rs.Next() {
		var s Datasource
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.Name,
			&s.Description,
			&s.Kind,
			&s.Configuration,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanDataset(r *sql.Row) (Dataset, error) {
	var s Dataset
	if err := r.Scan(
		&s.Id,
		&s.DatasourceId,
		&s.Name,
		&s.Description,
		&s.FrameName,
		&s.ResponseColumnName,
		&s.Properties,
		&s.PropertiesVersion,
		&s.Created,
	); err != nil {
		return Dataset{}, err
	}
	return s, nil
}

func ScanDatasets(rs *sql.Rows) ([]Dataset, error) {
	structs := make([]Dataset, 0, 16)
	var err error
	for rs.Next() {
		var s Dataset
		if err = rs.Scan(
			&s.Id,
			&s.DatasourceId,
			&s.Name,
			&s.Description,
			&s.FrameName,
			&s.ResponseColumnName,
			&s.Properties,
			&s.PropertiesVersion,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanModel(r *sql.Row) (Model, error) {
	var s Model
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.TrainingDatasetId,
		&s.ValidationDatasetId,
		&s.Name,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumnName,
		&s.LogicalName,
		&s.Location,
		&s.MaxRunTime,
		&s.Metrics,
		&s.MetricsVersion,
		&s.Created,
		&s.LabelId,
		&s.LabelName,
	); err != nil {
		return Model{}, err
	}
	return s, nil
}

func ScanModels(rs *sql.Rows) ([]Model, error) {
	structs := make([]Model, 0, 16)
	var err error
	for rs.Next() {
		var s Model
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.TrainingDatasetId,
			&s.ValidationDatasetId,
			&s.Name,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumnName,
			&s.LogicalName,
			&s.Location,
			&s.MaxRunTime,
			&s.Metrics,
			&s.MetricsVersion,
			&s.Created,
			&s.LabelId,
			&s.LabelName,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanBinomialModel(r *sql.Row) (BinomialModel, error) {
	var s BinomialModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.TrainingDatasetId,
		&s.ValidationDatasetId,
		&s.Name,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumnName,
		&s.LogicalName,
		&s.Location,
		&s.MaxRunTime,
		&s.Metrics,
		&s.MetricsVersion,
		&s.Created,
		&s.LabelId,
		&s.LabelName,
		&s.Mse,
		&s.RSquared,
		&s.Logloss,
		&s.Auc,
		&s.Gini,
	); err != nil {
		return BinomialModel{}, err
	}
	return s, nil
}

func ScanBinomialModels(rs *sql.Rows) ([]BinomialModel, error) {
	structs := make([]BinomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s BinomialModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.TrainingDatasetId,
			&s.ValidationDatasetId,
			&s.Name,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumnName,
			&s.LogicalName,
			&s.Location,
			&s.MaxRunTime,
			&s.Metrics,
			&s.MetricsVersion,
			&s.Created,
			&s.LabelId,
			&s.LabelName,
			&s.Mse,
			&s.RSquared,
			&s.Logloss,
			&s.Auc,
			&s.Gini,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanMultinomialModel(r *sql.Row) (MultinomialModel, error) {
	var s MultinomialModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.TrainingDatasetId,
		&s.ValidationDatasetId,
		&s.Name,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumnName,
		&s.LogicalName,
		&s.Location,
		&s.MaxRunTime,
		&s.Metrics,
		&s.MetricsVersion,
		&s.Created,
		&s.LabelId,
		&s.LabelName,
		&s.Mse,
		&s.RSquared,
		&s.Logloss,
	); err != nil {
		return MultinomialModel{}, err
	}
	return s, nil
}

func ScanMultinomialModels(rs *sql.Rows) ([]MultinomialModel, error) {
	structs := make([]MultinomialModel, 0, 16)
	var err error
	for rs.Next() {
		var s MultinomialModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.TrainingDatasetId,
			&s.ValidationDatasetId,
			&s.Name,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumnName,
			&s.LogicalName,
			&s.Location,
			&s.MaxRunTime,
			&s.Metrics,
			&s.MetricsVersion,
			&s.Created,
			&s.LabelId,
			&s.LabelName,
			&s.Mse,
			&s.RSquared,
			&s.Logloss,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanRegressionModel(r *sql.Row) (RegressionModel, error) {
	var s RegressionModel
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.TrainingDatasetId,
		&s.ValidationDatasetId,
		&s.Name,
		&s.ClusterName,
		&s.ModelKey,
		&s.Algorithm,
		&s.ModelCategory,
		&s.DatasetName,
		&s.ResponseColumnName,
		&s.LogicalName,
		&s.Location,
		&s.MaxRunTime,
		&s.Metrics,
		&s.MetricsVersion,
		&s.Created,
		&s.LabelId,
		&s.LabelName,
		&s.Mse,
		&s.RSquared,
		&s.MeanResidualDeviance,
	); err != nil {
		return RegressionModel{}, err
	}
	return s, nil
}

func ScanRegressionModels(rs *sql.Rows) ([]RegressionModel, error) {
	structs := make([]RegressionModel, 0, 16)
	var err error
	for rs.Next() {
		var s RegressionModel
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.TrainingDatasetId,
			&s.ValidationDatasetId,
			&s.Name,
			&s.ClusterName,
			&s.ModelKey,
			&s.Algorithm,
			&s.ModelCategory,
			&s.DatasetName,
			&s.ResponseColumnName,
			&s.LogicalName,
			&s.Location,
			&s.MaxRunTime,
			&s.Metrics,
			&s.MetricsVersion,
			&s.Created,
			&s.LabelId,
			&s.LabelName,
			&s.Mse,
			&s.RSquared,
			&s.MeanResidualDeviance,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanLabel(r *sql.Row) (Label, error) {
	var s Label
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.ModelId,
		&s.Name,
		&s.Description,
		&s.Created,
	); err != nil {
		return Label{}, err
	}
	return s, nil
}

func ScanLabels(rs *sql.Rows) ([]Label, error) {
	structs := make([]Label, 0, 16)
	var err error
	for rs.Next() {
		var s Label
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.ModelId,
			&s.Name,
			&s.Description,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}

func ScanService(r *sql.Row) (Service, error) {
	var s Service
	if err := r.Scan(
		&s.Id,
		&s.ProjectId,
		&s.ModelId,
		&s.Name,
		&s.Address,
		&s.Port,
		&s.ProcessId,
		&s.State,
		&s.Created,
	); err != nil {
		return Service{}, err
	}
	return s, nil
}

func ScanServices(rs *sql.Rows) ([]Service, error) {
	structs := make([]Service, 0, 16)
	var err error
	for rs.Next() {
		var s Service
		if err = rs.Scan(
			&s.Id,
			&s.ProjectId,
			&s.ModelId,
			&s.Name,
			&s.Address,
			&s.Port,
			&s.ProcessId,
			&s.State,
			&s.Created,
		); err != nil {
			return nil, err
		}
		structs = append(structs, s)
	}
	if err = rs.Err(); err != nil {
		return nil, err
	}
	return structs, nil
}
