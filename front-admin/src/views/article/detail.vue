<template>
  <div class="app-container">
    <div>
      <el-button  style="float: right;margin-right: 20px" size="small" @click="resetForm('ruleForm')">返回</el-button>
    </div>
    <div style="clear: both"></div>
    <el-form  style="margin-top: 20px;width: 1000px" size="small"  :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="文章名称" prop="article_name">
        <el-input v-model="ruleForm.article_name"></el-input>
      </el-form-item>
      <el-form-item label="文章分类" prop="category_id">
        <el-select v-model="ruleForm.category_id" placeholder="请选择文章分类">
          <el-option v-for="category in categories" :label="category.category_name" :value="category.id"></el-option>
        </el-select>
      </el-form-item>

      <el-form-item label="文章封面" prop="cover">
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
      <el-form-item label="文章标签" prop="tag_list">
        <el-tag
          :key="tag"
          v-for="tag in ruleForm.tag_list"
          closable
          :disable-transitions="false"
          @close="handleClose(tag)">
          {{tag}}
        </el-tag>
        <el-input
          class="input-new-tag"
          v-if="inputVisible"
          v-model="inputValue"
          ref="saveTagInput"
          size="small"
          @keyup.enter.native="handleInputConfirm"
          @blur="handleInputConfirm"
        >
        </el-input>
        <el-button v-else class="button-new-tag" size="small" @click="showInput">+</el-button>
      </el-form-item>


      <el-form-item label="文章简介" prop="description" width="200">
        <el-input type="textarea" v-model="ruleForm.description"></el-input>
      </el-form-item>
      <el-form-item label="文章内容" prop="content" width="200">
        <mavon-editor ref=content @imgAdd="$imgAdd"  :toolbars="editConfig" v-model="ruleForm.content"/>
      </el-form-item>

      <el-form-item label="排序" prop="sort">
        <el-input v-model="ruleForm.sort"></el-input>
      </el-form-item>
      <el-form-item label="显示">
        <el-switch v-model="ruleForm.is_show" active-value="1" inactive-value="2"></el-switch>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">保存</el-button>
        <el-button @click="resetForm('ruleForm')">返回</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
    import { uploadImage,saveArticle,shoeArticle } from '@/api/article'
    export default {
        props :["isEdit","articleId","categories"],
        data() {
            return {

                inputVisible: false,
                inputValue: '',
                ruleForm: {
                    article_name: '',
                    category_id: '',
                    type: [],
                    tag_list: [],
                    resource: '',
                    description: '',
                    content: '',
                    is_show: '1',
                    sort: 1,
                    cover:""
                },
                rules: {
                    article_name: [
                        { required: true, message: '请输入文章名称', trigger: 'blur' },
                    ],
                    category_id: [
                        { required: true, message: '请选择文章分类', trigger: 'change' }
                    ],
                    description: [
                        { required: true, message: '请填写简介', trigger: 'blur' }
                    ],
                    content: [
                        { required: true, message: '请填写内容', trigger: 'blur' }
                    ]
                },
                editConfig:{
                    bold: true, // 粗体
                    italic: true, // 斜体
                    header: true, // 标题
                    underline: true, // 下划线
                    strikethrough: true, // 中划线
                    mark: true, // 标记
                    superscript: true, // 上角标
                    subscript: true, // 下角标
                    quote: true, // 引用
                    ol: true, // 有序列表
                    ul: true, // 无序列表
                    link: true, // 链接
                    imagelink: true, // 图片链接
                    code: true, // code
                    table: true, // 表格
                    fullscreen: true, // 全屏编辑
                    /* 1.4.2 */
                    navigation: true, // 导航目录
                    /* 2.1.8 */
                    alignleft: true, // 左对齐
                    aligncenter: true, // 居中
                    alignright: true, // 右对齐
                    /* 2.2.1 */
                    subfield: true, // 单双栏模式
                    preview: true, // 预览
                }
            };
        },
        methods: {
            submitForm(formName) {
                this.$refs[formName].validate((valid) => {
                    if (valid) {
                        let params = this.ruleForm;
                        if(this.articleId>0){
                            params.id = Number(this.articleId)
                        }
                        params.sort = Number( params.sort);
                        params.is_show = Number( params.is_show);
                        saveArticle(params).then((res)=>{
                            if (res.code===0){
                                this.$message.success("保存成功");
                                this.resetForm('ruleForm')
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
            resetForm(formName) {
                this.$refs[formName].resetFields();
                this.$emit('detail-hide',false);
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
            handleClose(tag) {
                this.ruleForm.tag_list.splice(this.ruleForm.tag_list.indexOf(tag), 1);
            },

            showInput() {
                this.inputVisible = true;
                this.$nextTick(_ => {
                    this.$refs.saveTagInput.$refs.input.focus();
                });
            },

            handleInputConfirm() {
                let inputValue = this.inputValue;
                if (inputValue) {
                    this.ruleForm.tag_list.push(inputValue);
                }
                this.inputVisible = false;
                this.inputValue = '';
            },
            $imgAdd(pos, $file){
                var formdata = new FormData();
                formdata.append('image', $file);
                uploadImage(formdata).then((res) => {
                    if (res.code===0){
                        this.$refs.content.$img2Url(pos, res.data.file);
                    }else{
                        this.$message.error(res.message);
                    }

                })
            }
        },
        created() {
          if (this.articleId>0){
              shoeArticle({article_id: this.articleId}).then((res)=>{
                  if (res.code===0){
                      this.ruleForm = res.data;
                      this.ruleForm.is_show=this.ruleForm.is_show.toString()
                  }
              })
          }
        }
    }
</script>
