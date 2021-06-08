
#接口名称：getRarity
接口描述：通过输入士兵id获取稀有度
输入参数：字段         类型               说明
        id        string              士兵id
        Umap     map[string]*Info     存储士兵数据的数据结构
输出参数： 字段        类型               说明 
        rarity     string           士兵的稀有度

#接口名称：getCombatPoints
接口描述：通过输入士兵id获取战力
输入参数：字段         类型               说明
        id        string              士兵id
        Umap     map[string]*Info     存储士兵数据的数据结构
输出参数： 字段        类型                说明
        CombatPoints     string         士兵的战力

#接口名称：getSoldiersJson
接口描述：通过阶段来获取已解锁的士兵信息
输入参数：字段         类型               说明
        stage        string             阶段
        Umap     map[string]*Info     存储士兵数据的数据结构
输出参数： 字段        类型                说明
        data       []*Info            存储符合要求的士兵数据的数据结构

#接口名称：getArmyAssembled
接口描述：通过稀有度，阶段来获取已解锁的所有士兵的信息
输入参数：字段         类型               说明
        rarity        string          稀有度
        stage         string          阶段
        Umap     map[string]*Info     存储士兵数据的数据结构
输出参数： 字段        类型                说明
data       []*Info              存储符合要求的士兵数据的数据结构