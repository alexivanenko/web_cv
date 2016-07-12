
function attachRemoveEvent() {
    $('.close').click(function() {
        var id = $(this).attr('id');

        $.ajax({
            url: "/admin/remove_project",
            dataType: 'json',
            data: {'id': id},
            beforeSend: function ( xhr ) {
                return confirm('Do you really want to delete this project?');
            },

            error:function(jqXHR, textStatus, errorThrown) {
                alert("Error Occurred :: "+textStatus+";\r"+errorThrown+"\r"+jqXHR.responseText);
            },

            success: function ( data ) {
                $('#item_' + data.id).remove();
            }
        });
    });
}

function attachEditEvent() {
    $('.image_portfolio').click(function() {
        var projectId = $(this).find('button').attr('id');
        var portfolioId = $('input[name="id"]').val();

        $.ajax({
            url: "/admin/get_project_data",
            dataType: 'json',
            data: {'id': portfolioId, 'project_id': projectId},

            error:function(jqXHR, textStatus, errorThrown) {
                alert("Error Occurred :: "+textStatus+";\r"+errorThrown+"\r"+jqXHR.responseText);
            },

            success: function ( data ) {
                if (data.name != "") {
                    $('input[name="projectId"]').val(projectId);
                    $('input[name="projectImage"]').val(data.image);
                    $('input[name="Name"]').val(data.name);
                    $('select[name="Category"]').val(data.category);
                    $('input[name="Url"]').val(data.url);
                    $('input[name="ShortDescription"]').val(data.short_description);
                    $('textarea[name="Description"]').text(data.description);
                    $('input[name="Order"]').val(data.order);
                }
            }
        });

        return false;
    });
}

function attachResetEvent() {
    $('#resetButton').click(function() {
        $('textarea[name="Description"]').text('');
        $('input[name="projectId"]').val('');
        $('input[name="projectImage"]').val('');
        $("#projectForm")[0].reset();
    });
}