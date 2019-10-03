#!/usr/bin/env bash

set -e
set -u
set -o pipefail

if [ -n "${PARAMETER_STORE:-}" ]; then
  export EVALUACION_INSCRIPCION_CRUD__PGUSER="$(aws ssm get-parameter --name /${PARAMETER_STORE}/evaluacion_inscripcion_crud/db/username --output text --query Parameter.Value)"
<<<<<<< HEAD
  export EVALUACION_INSCRIPCION_CRUD__PGUSER="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/evaluacion_inscripcion_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
=======
  export EVALUACION_INSCRIPCION_CRUD__PGPASS="$(aws ssm get-parameter --with-decryption --name /${PARAMETER_STORE}/evaluacion_inscripcion_crud/db/password --output text --query Parameter.Value)"
fi

exec ./main "$@"
>>>>>>> ca608151a1e7a4fa6c6a5d6d009429feb4a0c5e0
