<template>
  <div class="app-container">

    <div style="clear: both"></div>
    <el-form  style="margin-top: 20px;width: 1000px" size="small"  :model="ruleForm"  ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="站点名称" prop="title">
        <el-input v-model="ruleForm.title"></el-input>
      </el-form-item>
      <el-form-item label="关键词" prop="keyword">
        <el-input v-model="ruleForm.keyword"></el-input>
      </el-form-item>
      <el-form-item label="站点简介" prop="description">
        <el-input type="textarea" v-model="ruleForm.description"></el-input>
      </el-form-item>

      <el-form-item label="微信二维码" prop="cover">
        <el-upload
          class="avatar-uploader"
          name="image"
          action="/api/backend/upload/image"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload">
          <img v-if="ruleForm.cover" :src="ruleForm.cover" class="avatar">
          <i v-else class="el-icon-plus avatar-uploader-icon"></i>
        </el-upload>
      </el-form-item>

      <el-form-item label="统计代码" prop="code" width="200">
        <el-input type="textarea" v-model="ruleForm.code"></el-input>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">保存</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
    import { uploadImage } from '@/api/article'
    import { siteInfo,saveSite } from '@/api/site'
    export default {
        data() {
            return {
                ruleForm: {
                    title: '',
                    keyword: '',
                    description: '',
                    cover:"",
                    code:"",
                },
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        let params = this.ruleForm;
                        saveSite(params).then((res)=>{
                            if (res.code===0){
                                this.$message.success("保存成功");
                            }else {
                                this.$message.error(res.message)
                            }
                        })
                    } else {
                        console.log('error submit!!');
                        return false;
                    }
                });
            },

            handleAvatarSuccess(res, file) {
                if (res.code===0){
                    this.ruleForm.cover = res.data.file;
                }else{
                    this.$message.error(res.message);
                }
            },
            beforeAvatarUpload(file) {
                const isJPG = file.type === 'image/jpeg'||file.type === 'image/png';
                const isLt2M = file.size / 1024 / 1024 < 5;

                if (!isJPG) {
                    this.$message.error('图片只能是 JPG、PNG 格式!');
                }
                if (!isLt2M) {
                    this.$message.error('图片大小不能超过 5MB!');
                }
                return isJPG && isLt2M;
            },

        },
        created() {
            siteInfo().then((res)=>{
                if (res.code===0){
                    this.ruleForm = res.data;
                }
            })
        }
    }
</script>
