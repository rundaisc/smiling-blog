<template>
  <div class="app-container">
    <div>
      <el-form size="small">
        <el-form-item style="float: right">
          <el-button type="primary" @click="goCreate">新建</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      size="small"
      v-loading="listLoading"
      :data="list"
      element-loading-text="Loading"
    >
      <el-table-column align="center" label="名称" prop="link_name" >
      </el-table-column>
      <el-table-column align="center" label="地址" prop="url" >
      </el-table-column>
      <el-table-column align="center" label="排序" prop="sort" >
      </el-table-column>
      <el-table-column class-name="status-col" label="显示"  align="center">
        <template slot-scope="scope">
          <el-tag v-if="scope.row.is_show==1">显示</el-tag>
          <el-tag v-if="scope.row.is_show==2" type="info">隐藏</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        align="center"
        width="100">
        <template slot-scope="scope" >
          <el-button @click="handleClick(scope.row)" type="text" size="small">编辑</el-button>
          <el-button type="text" @click="delClick(scope.row)" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>


    <el-dialog title="友情链接" width="600px" ref="form" :model="form" :visible.sync="dialogFormVisible">
      <el-form :model="form" label-width="100px">
        <el-form-item label="名称" label-width="100px">
          <el-input v-model="form.link_name" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="URL" label-width="100px">
          <el-input  v-model="form.url" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="排序" label-width="100px">
          <el-input v-model="form.sort" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="显示" label-width="100px">
          <el-switch v-model="form.is_show"
                     :active-value="1"
                     :inactive-value="2"></el-switch>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="saveCategory">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { saveFlink,getList,deleteFlink } from '@/api/flink'
export default {
  data() {
    return {
      list: null,
      listLoading: false,
      currentPage:1,
        dialogFormVisible:false,
        form: {
            link_name: '',
            is_show: 1,
            url: "",
            sort: 99,
            id:0,
        },
        categories:[]
    }
  },
  created() {
    this.fetchData();
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      getList().then(response => {
        this.list = response.data;
        this.listLoading = false
      })
    },

      goCreate(){
        this.dialogFormVisible = true;
        this.form = {
            flink_name: '',
            is_show: 1,
            url: "",
            sort: 99,
            id:0,
        }
      },
      saveCategory(){
        let that = this;
        let params = that.form;
          params.is_show = Number(params.is_show);
          params.sort = Number(params.sort);
          saveFlink(params).then(res=>{
              if (res.code==0){
                  that.$message.success("保存成功");
                  this.dialogFormVisible = false;
                  that.fetchData()
              }else{
                  that.$message.error(res.message)
              }
          });
      },
      handleClick(row){
        let params = row;
        this.form = params;
        this.dialogFormVisible = true;
      },
      delClick(row){
        let that = this;
          this.$confirm('确定删除该友链?', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning'
          }).then(() => {
              let params = {
                  id:row.id
              };
              deleteFlink(params).then((res)=>{
                  that.$message.success("保存成功")
                  that.fetchData()
              })
          }).catch(() => {

          });
      }

  }
}
</script>
