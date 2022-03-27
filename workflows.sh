# generate workflow files for all services
# https://dev.to/stackdumper/setting-up-ci-for-microservices-in-monorepo-using-github-actions-5do2

# read the workflow template
WORKFLOW_TEMPLATE=$(cat .github/workflow-template.yaml)

# iterate each svc in service directory
for SVC in $(ls service); do
    echo "generating workflow for service/${SVC}/api"
    API_WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{NAME}}/${SVC}/g" | sed "s/{{TYPE}}/api/g")
    echo "${API_WORKFLOW}" > .github/workflows/${SVC}-api.yaml

    echo "generating workflow for service/${SVC}/rpc"
    RPC_WORKFLOW=$(echo "${WORKFLOW_TEMPLATE}" | sed "s/{{NAME}}/${SVC}/g" | sed "s/{{TYPE}}/rpc/g")
    echo "${RPC_WORKFLOW}" > .github/workflows/${SVC}-rpc.yaml
done
