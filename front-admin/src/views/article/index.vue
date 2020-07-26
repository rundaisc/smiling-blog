<template>
  <div class="app-container" v-if="!isEdit">
    <div>
      <el-form  size="small"  :inline="true" :model="form" class="demo-form-inline">
        <el-form-item label="文章名称">
          <el-input v-model="form.article_name"  placeholder="文章名称"></el-input>
        </el-form-item>
        <el-form-item label="分类">
          <el-select clearable v-model="form.category_id" placeholder="分类">
            <el-option v-for="category in categories" :label="category.category_name" :value="category.id"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="显示">
          <el-select clearable  v-model="form.is_show" placeholder="显示">
            <el-option label="是" value="1"></el-option>
            <el-option label="否" value="2"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="goSearch">查询</el-button>
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
      <el-table-column align="center" label="ID" prop="id" >
      </el-table-column>
      <el-table-column align="center" label="名称" prop="article_name" >
      </el-table-column>
      <el-table-column align="center" label="浏览量" prop="view" >
      </el-table-column>
      <el-table-column align="center" label="评论数" prop="comment" >
      </el-table-column>
      <el-table-column label="分类" prop="category_name"  align="center">
      </el-table-column>
      <el-table-column align="center" prop="created_at" label="创建时间" width="180">
      </el-table-column>
      <el-table-column class-name="status-col" label="显示"  align="center">
        <template slot-scope="scope">
          <el-tag type="success" v-if="scope.row.is_show===1">显示</el-tag>
          <el-tag type="info" v-if="scope.row.is_show!==1">隐藏</el-tag>
        </template>
      </el-table-column>
      <el-table-column
        fixed="right"
        label="操作"
        width="100">
        <template slot-scope="scope">
          <el-button @click="goEdit(scope.row)" type="text" size="small">编辑</el-button>
          <el-button type="text"   @click="goDelete(scope.row)" size="small">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      class="page-box"
      :current-page.sync="currentPage"
      layout="total, prev, pager, next"
      :total="total">
    </el-pagination>
  </div>
  <detail-page :isEdit="isEdit" :categories="categories" :articleId="articleId" @detail-hide="detailBack" v-else-if="isEdit"></detail-page>
</template>

<script>
import { getList,deleteArticle } from '@/api/article'
import { getList as categoryListApi } from '@/api/category'
import detailPage from './detail'
export default {
    components: {
        detailPage,
    },
  data() {
    return {
      list: null,
      listLoading: true,
      currentPage:1,
        total:0,
        isEdit:false,
        articleId : 0,
        categories:[],
        form:{
            article_name:"",
            category_id:"",
            is_show:"",
        }
    }
  },
  created() {
      this.fetchData();
      categoryListApi().then((res)=>{
          this.categories = res.data;
      })
  },
  methods: {
    fetchData() {
      this.listLoading = true;
      let params = {
          article_name:this.form.article_name,
          category_id:Number(this.form.category_id),
          is_show:Number(this.form.is_show),
          page:Number(this.currentPage),
          page_size:10,
      };
      getList(params).then(response => {
        this.list = response.data.list;
        this.total =response.data.total;
        this.listLoading = false
      })
    },

      goCreate(){
        this.isEdit = true;
        this.articleId = 0;
      },
      goSearch(){
         this.currentPage = 1;
         this.fetchData()
      },
      goEdit(row){
          this.isEdit = true;
          this.articleId = row.id;
      },
      detailBack(){
        this.isEdit = false;
        this.currentPage =1 ;
          this.fetchData();
      },
      goDelete(row){
        let that = this;
          this.$confirm('你确定删除该文章？', '提示', {
              confirmButtonText: '确定',
              cancelButtonText: '取消',
              type: 'warning',}).then(()=>{
              deleteArticle({id:row.id}).then((res)=>{
                  if (res.code===0){
                      that.$message.success("删除成功");
                      that.currentPage =1 ;
                      that.fetchData();
                  }else{
                      that.$message.error(res.message)
                  }
              })
          })
      }

  }
}
</script>
