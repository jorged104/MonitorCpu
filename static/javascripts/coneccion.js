$( document ).ready(function() {


    var tablaCuerpo = $("#cuerpo_tabla");
    var mensajec = $("#mensaje");

    $.ajax({
        url: 'panel',
        dataType: 'json',
        type: 'get',
        contentType: 'application/x-www-form-urlencoded',
        data: 'hola',
        success: function( data, textStatus, jQxhr ){
         var tabla = data.tabla;
         var datos = '';
         var m = '';
         tabla.forEach(element => {
             datos  += `<tr><td>${element.pid}</td><td>${element.usuario}</td><td>${element.estado}</td><td>${element.ram}</td><td>${element.nombre}</td><td> <button class="btn btn-primary" onclick="kill(${element.pid})">Kill</button> </td></tr>`
         });
         tablaCuerpo.html(datos);
         m += `Procesos Totales :  ${data.totales} 
         Procesos Ejecutando : ${data.ejecutando} 
          Processo suspendidos : ${data.suspendido} Procesos detenidos ${data.detenidos}`
         mensajec.html(m); 
        }

    });
});


function kill(pid)
{
    $.ajax({
        url: 'panel/matar',
        dataType: 'json',
        type: 'get',
        contentType: 'application/x-www-form-urlencoded',
        data: {valor : pid},
        success: function( data, textStatus, jQxhr ){
            alert(" Se elimino " + pid )
        }

    });
}