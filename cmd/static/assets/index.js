$(function () {
    $(".add-file").on("click", function () {
        $("#file-input").click();
    });

    $("#file-input").on("change", function(){
        let files = $('#file-input')[0].files;
        if (files.length <= 0) {
            console.log('请选择文件后在上传')
            return
        }

        let fd = new FormData();
        fd.append('file', files[0]);
        let fileFullName = $('#file-input').val();
        let fileArr = fileFullName.split("\\")
        $("#filename").text(fileArr[fileArr.length - 1]);
        $.ajax({
            method: 'post',
            url: '/upload',
            data: fd,
            processData: false,
            contentType: false,
            success: function (res) {
                
                // console.log(res);
            },
            xhr: function(){
                var xhr = new XMLHttpRequest(); 
            
                xhr.upload.addEventListener('progress', function(e){
                
                var vnum = Math.floor((e.loaded / e.total) * 100);
                //var progressRate = vnum + '%'; 
                $('#progressvens').attr('value', vnum);
              })
              return xhr; 
            }
        });
        $('#file-input').val("");
    });

});