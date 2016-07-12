
function attachAddRowsEvents() {
    $(".btn.icon-btn.btn-success.service").click(function() {
        $("#servicesRows").append(getServiceHtml(getLastIndex("service")));
    });
    $(".btn.icon-btn.btn-success.testimonial").click(function() {
        $("#testimonialsRows").append(getTestimonialHtml(getLastIndex("testimonial")));
    });
}

function getLastIndex(type) {
    var lastNameId = $(".form-control."+type+"-name").last().attr('id');
    var index = 0;

    if (typeof  lastNameId != 'undefined') {
        index = parseInt(lastNameId.replace(/[^0-9]/gi, ''));
        if ($.isNumeric(index))
            index ++;
    }

    return index;
}

function getServiceHtml(index) {
    return '<div class="form-inline text-center"> \
                <div class="form-group"> \
                    <label for="inputServiceName'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control service-name" id="inputServiceName'+index+'" name="Service.Name" placeholder="SERVICE" value=""> \
                    </div> \
                </div> \
                <div class="form-group"> \
                    <label for="inputServiceDescription'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control" id="inputServiceDescription'+index+'" name="Service.Description" placeholder="Description" value=""> \
                    </div> \
                </div> \
                <div class="form-group"> \
                    <label for="inputServiceIcon'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control" id="inputServiceIcon'+index+'" name="Service.Icon" placeholder="Icon Class" value=""> \
                    </div> \
                </div> \
                ' +getRemoveBtnHtml()+ ' \
            </div>\
        ';
}

function getTestimonialHtml(index) {
    return '<div class="form-inline text-center"> \
                <div class="form-group"> \
                    <label for="inputTestimonialContent'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control testimonial-name" id="inputTestimonialContent'+index+'" name="Testimonial.Content" placeholder="Good Job!" value=""> \
                    </div> \
                </div> \
                <div class="form-group"> \
                    <label for="inputTestimonialWriterName'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control" id="inputTestimonialWriterName'+index+'" name="Testimonial.WriterName" placeholder="Sergey Brin" value=""> \
                    </div> \
                </div> \
                <div class="form-group"> \
                    <label for="inputTestimonialWriterCompany'+index+'"></label> \
                    <div class="col-sm-2"> \
                        <input type="text" class="form-control service-name" id="inputTestimonialWriterCompany'+index+'" name="Testimonial.WriterCompany" placeholder="Google" value=""> \
                    </div> \
                </div> \
                ' +getRemoveBtnHtml()+ ' \
            </div> \
        ';
}

function getRemoveBtnHtml() {
    return '<div class="form-group"> \
                <div class="col-sm-offset-2 col-sm-5"> \
                    <a class="btn icon-btn btn-danger" href="#" onclick="removeRow(this); return false;"> \
                        <span class="glyphicon btn-glyphicon glyphicon-trash img-circle text-danger"></span> \
                        Delete \
                    </a> \
                </div> \
            </div> \
        ';
}

function removeRow(_button) {
    $(_button).parent().parent().parent().remove();
}
