import 'dart:async';
import 'dart:math';

void main() async {
  print("Cargando clima...");

  final climaStream = obtenerClima(5);
  await for (var clima in climaStream) {
    print(clima);

    var temperatura = extraerTemperatura(clima);
    String resultado = await procesarTemperatura(temperatura);
    print("Resultado del analisis: $resultado\n");
  }
}

int extraerTemperatura(String clima) {
  try {
    return int.parse(clima.split("Temperatura: ")[1].split("°")[0].trim());
  } catch (e) {
    print("Error al extraer la temperatura: $e");
    return 0;
  }
}

Stream<String> obtenerClima(int limite) async* {
  final lugares = ["Arabia", "Gabon", "Brasil", "Kenia"];
  final condiciones = ["Solazo", "Palo de agua", "Nublado", "Mucha brisa"];
  final random = Random();
  int contador = 0;

  while (contador < limite) {
    await Future.delayed(Duration(seconds: 2));
    String lugar = lugares[random.nextInt(lugares.length)];
    String clima = condiciones[random.nextInt(condiciones.length)];
    int temperatura = random.nextInt(35);

    yield "Lugar: $lugar, Clima: $clima, Temperatura: $temperatura°C";
    contador++;
  }
}

Future<String> procesarTemperatura(int temperatura) async {
  await Future.delayed(Duration(seconds: 1));
  if (temperatura < 10) {
    return "Merida ❄️";
  } else if (temperatura < 25) {
    return "Apure 🌤️";
  } else {
    return "Zulia 🔥";
  }
}
